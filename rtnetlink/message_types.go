package rtnetlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "github.com/vishvananda/go-netlink"

// netlink.MessageType's for RTNETLINK
const (
	RTM_NEWLINK netlink.MessageType = iota + netlink.MIN_TYPE
	RTM_DELLINK
	RTM_GETLINK
	RTM_SETLINK
	RTM_NEWADDR
	RTM_DELADDR
	RTM_GETADDR
	RTM_RESERVED23
	RTM_NEWROUTE
	RTM_DELROUTE
	RTM_GETROUTE
	RTM_RESERVED27
	RTM_NEWNEIGH
	RTM_DELNEIGH
	RTM_GETNEIGH
	RTM_RESERVED31
	RTM_NEWRULE
	RTM_DELRULE
	RTM_GETRULE
	RTM_RESERVED35
	RTM_NEWQDISC
	RTM_DELQDISC
	RTM_GETQDISC
	RTM_RESERVED39
	RTM_NEWTCLASS
	RTM_DELTCLASS
	RTM_GETTCLASS
	RTM_RESERVED43
	RTM_NEWTFILTER
	RTM_DELTFILTER
	RTM_GETTFILTER
	RTM_RESERVED47
	RTM_NEWACTION
	RTM_DELACTION
	RTM_GETACTION
	RTM_RESERVED51
	RTM_NEWPREFIX
	RTM_RESERVED53
	RTM_RESERVED54
	RTM_RESERVED55
	RTM_RESERVED56
	RTM_RESERVED57
	RTM_GETMULTICAST
	RTM_RESERVED59
	RTM_RESERVED60
	RTM_RESERVED61
	RTM_GETANYCAST
	RTM_RESERVED63
	RTM_NEWNEIGHTBL
	RTM_RESERVED65
	RTM_GETNEIGHTBL
	RTM_SETNEIGHTBL
	RTM_NEWNDUSEROPT
	RTM_RESERVED69
	RTM_RESERVED70
	RTM_RESERVED71
	RTM_NEWADDRLABEL
	RTM_DELADDRLABEL
	RTM_GETADDRLABEL
	RTM_RESERVED75
	RTM_RESERVED76
	RTM_RESERVED77
	RTM_GETDCB
	RTM_SETDCB
	RTM_BASE = RTM_NEWLINK
	RTM_MAX  = RTM_SETDCB
)
