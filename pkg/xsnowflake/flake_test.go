package xsnowflake

import (
	"github.com/sony/sonyflake"
	"github.com/sony/sonyflake/awsutil"
	"math"
	"testing"
	"time"
)

func Test_Flake(t *testing.T) {

	sf = NewFlake(NewConfig())

	id, err := sf.NextID()
	// 132800044205082596
	// 18位

	tn := time.Now().Unix()
	// 1656991794
	// 10位

	res := sonyflake.Decompose(id)

	mid, _ := awsutil.AmazonEC2MachineID()
	var ui64 uint16 = math.MaxUint16

	t.Log("DONE", id, err, res, tn, mid, ui64)
}
