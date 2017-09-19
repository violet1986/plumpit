package collator

import (
	"log"
	"plumpit/protos"
)

func PrintSender(msg protos.PitMessage) error {
	switch msg.PitType {
	case protos.EnumPitType_QUERY_INFO:
		q := msg.GetQueryInfo()
		log.Println(q.GetQueryId(), q.GetStatus())
	case protos.EnumPitType_EXEC_INFO:
		q := msg.GetExecInfo()
		nodekey := q.GetNodeKey()
		log.Println(nodekey.SegId, nodekey.ProcId, nodekey.NodeId, q.GetStatus(), q.GetPlanRows(), q.GetNodeType())
	case protos.EnumPitType_PROCESS_INFO:
		q := msg.GetProcInfo()
		insight := q.GetInsight()
		for _, instr := range insight {
			nodekey := instr.GetKey()
			log.Println("\t", nodekey.SegId, nodekey.ProcId, nodekey.NodeId, instr.GetTupleCount())
		}
	}
	return nil
}
