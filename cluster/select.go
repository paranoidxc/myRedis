package cluster

import "myredis/interface/resp"

func execSelect(cluster *ClusterDatabase, c resp.Connection, cmdArgs [][]byte) resp.Reply {
	return cluster.ClusterExec(c, cmdArgs)
	//return cluster.db.Exec(c, cmdArgs)
}
