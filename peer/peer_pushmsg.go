package peer

import (
	"github.com/michain/dotcoin/protocol"
	"math/rand"
	"github.com/michain/dotcoin/logx"
)

func (p *Peer) PushAddrMsg(addresses []string) error {
	addressCount := len(addresses)

	// Nothing to send.
	if addressCount == 0 {
		return nil
	}

	msg := protocol.NewMsgAddr()
	msg.AddrList = make([]string, addressCount)
	copy(msg.AddrList, addresses)

	// Randomize the addresses sent if there are more than the maximum allowed.
	if addressCount > protocol.MaxAddrPerMsg {
		// Shuffle the address list.
		for i := 0; i < protocol.MaxAddrPerMsg; i++ {
			j := i + rand.Intn(addressCount-i)
			msg.AddrList[i], msg.AddrList[j] = msg.AddrList[j], msg.AddrList[i]
		}

		// Truncate it to the maximum size.
		msg.AddrList = msg.AddrList[:protocol.MaxAddrPerMsg]
	}

	//set single send
	msg.SetNeedBroadcast(false)

	p.SendSingleMessage(msg)
	return nil
}

func (p *Peer) PushVersion(msg *protocol.MsgVersion) error{
	logx.DevPrintf("Peer.PushVersion %v", msg)
	if msg.AddrFrom == ""{
		msg.AddrFrom = p.GetSeedAddr()
	}
	p.SendSingleMessage(msg)
	return nil
}

func (p *Peer) PushGetBlocks(msg *protocol.MsgGetBlocks) error{
	logx.DevPrintf("Peer.PushGetBlocks peer:%v msg:%v", p.GetListenAddr(), msg)
	p.SendSingleMessage(msg)
	return nil
}

func (p *Peer) PushBlock(msg *protocol.MsgBlock) error{
	logx.DevPrintf("Peer.PushBlock peer:%v msg:%v", p.GetListenAddr(), msg)
	p.SendSingleMessage(msg)
	return nil
}