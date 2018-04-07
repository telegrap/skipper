// Package swarm implements functionality to exchange information
// between skipper instances. It is meant for low latency, weakly
// consistent data. As an example tere is the filter clusterRatelimit
// implementation, that uses swarm data exchnage to have a global
// state of current requests.
//
// Background information:
//
// The current skipper implementation uses
// hashicorp's memberlist, https://github.com/hashicorp/memberlist,
// which is an implementation of the swim protocol. You can find a
// detailed paper at
// http://www.cs.cornell.edu/~asdas/research/dsn02-SWIM.pdf.
//
// Quote from a nice overview https://prakhar.me/articles/swim/
//
//     The SWIM or the Scalable Weakly-consistent Infection-style process
//     group Membership protocol is a protocol used for maintaining
//     membership amongst processes in a distributed system.
//
//
package swarm