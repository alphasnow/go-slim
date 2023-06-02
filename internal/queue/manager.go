package queue

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"log"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

type Manager struct {
	RdsOpt    asynq.RedisClientOpt
	Conf      asynq.Config
	Tasks     *Tasks
	client    *asynq.Client
	inspector *asynq.Inspector
	server    *asynq.Server
}

func (m *Manager) newServe() (*asynq.Server, *asynq.ServeMux) {
	srv := asynq.NewServer(
		m.RdsOpt,
		m.Conf,
	)
	mux := asynq.NewServeMux()

	// mux.Handle(backup.EmailSendType, &backup.EmailSendTask{})
	// mux.HandleFunc(DemoTaskType, ProcessDemoTask)
	for t, h := range m.Tasks.GetHandlerMap() {
		mux.Handle(t, h)
	}
	return srv, mux
}

func (m *Manager) Start() *asynq.Server {

	srv, mux := m.newServe()

	//if err := srv.Run(mux); err != nil {
	//	log.Fatalf("could not run server: %v", err)
	//}
	if err := srv.Start(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

	m.server = srv
	return m.server
}

func (m *Manager) StartBlocking() *asynq.Server {

	srv, mux := m.newServe()

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

	m.server = srv
	return m.server
}

func (m *Manager) Close() {
	m.server.Shutdown()
}

func (m *Manager) NewInspector() *asynq.Inspector {
	inspector := asynq.NewInspector(m.RdsOpt)
	return inspector
}
func (m *Manager) Inspector() *asynq.Inspector {
	if m.inspector != nil {
		return m.inspector
	}
	m.inspector = asynq.NewInspector(m.RdsOpt)
	return m.inspector
}

func (m *Manager) NewClient() *asynq.Client {
	client := asynq.NewClient(m.RdsOpt)
	return client
}
func (m *Manager) Client() *asynq.Client {
	if m.client != nil {
		return m.client
	}
	m.client = asynq.NewClient(m.RdsOpt)
	return m.client
}

//	p := &tasks.EmailSendTask{}
//	p.EnqueueTask(client, tasks.EmailSendPayload{From: "aa@qq.com", To: "bb@qq.com"})

func (m *Manager) EnqueueTask(name string, p interface{}, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	task, err := asynq.NewTask(name, payload, opts...), nil
	if err != nil {
		return nil, err
	}

	info, err := m.Client().Enqueue(task)
	if err != nil {
		return nil, err
	}

	return info, nil
}
