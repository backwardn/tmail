package api

// !!! scope doit etre initialisé avant d'utiliser ce package

import (
	"github.com/toorop/tmail/core"
)

// USER

// UserGetByLoging returns an User by his login
func UserGetByLogin(login string) (user *core.User, err error) {
	return core.UserGetByLogin(login)
}

// UserAdd add a new usere
func UserAdd(login, passwd, mbQuota string, haveMailbox, authRelay bool) error {
	return core.UserAdd(login, passwd, mbQuota, haveMailbox, authRelay)
}

// UserDel delete an user (keep his mailboxe)
func UserDel(login string) error {
	return core.UserDel(login)
}

// UserList return all users
func UserGetAll() (users []core.User, err error) {
	return core.UserList()
}

// MAILBOXES
// MailboxAdd create a new mailbox
func MailboxAdd(mailbox string) error {
	return core.MailboxAdd(mailbox)
}

// MailboxDel delete a mailbox
func MailboxDel(mailbox string) error {
	return core.MailboxDel(mailbox)
}

// MailboxList return all mailboxes
func MailboxList() (mailboxes []core.Mailbox, err error) {
	return core.MailboxList()
}

// RELAY IP
// RelayIpAdd add an IP authozed to relay through tmail
func RelayIpAdd(ip string) error {
	return core.RelayIpAdd(ip)
}

// RelayIpDel remove an ip from authorized IP
func RelayIpDel(ip string) error {
	return core.RelayIpDel(ip)
}

// RelayIpGetAll returns all IPs which are authorized to relay through tmail
func RelayIpGetAll() (ips []core.RelayIpOk, err error) {
	return core.RelayIpGetAll()
}

// Queue
// QueueGetMessages returns all message in queue
func QueueGetMessages() ([]core.QMessage, error) {
	return core.QueueListMessages()
}

// QueueGetMessage return a message by its id
func QueueGetMessage(id int64) (core.QMessage, error) {
	return core.QueueGetMessageById(id)
}

// QueueDiscardMsgByKey discard a message (delete without bouncing) by his id
func QueueDiscardMsg(id int64) error {
	m, err := core.QueueGetMessageById(id)
	if err != nil {
		return err
	}
	return m.Discard()
}

// QueueBounceMsgByKey bounce a message by his key
func QueueBounceMsg(id int64) error {
	m, err := core.QueueGetMessageById(id)
	if err != nil {
		return err
	}
	return m.Bounce()
}

// ROUTES
// RoutesGet returns all routes
func RoutesGet() ([]core.Route, error) {
	return core.GetAllRoutes()
}

// RoutesAdd adds en new route
func RoutesAdd(host, localIp, remoteHost string, remotePort, priority int, user, mailFrom, smtpAuthLogin, smtpAuthPasswd string) error {
	return core.AddRoute(host, localIp, remoteHost, remotePort, priority, user, mailFrom, smtpAuthLogin, smtpAuthPasswd)
}

// RoutesDel delete route routeId
func RoutesDel(routeId int64) error {
	return core.DelRoute(routeId)
}

// RCPTHOSTS ie locals domains

// RcptHostAdd add a rcpthost
func RcpthostAdd(host string, isLocal bool) error {
	return core.RcpthostAdd(host, isLocal)
}

// RcpthostDel delete a rcpthost
func RcpthostDel(host string) error {
	return core.RcpthostDel(host)
}

// RcpthostList returns all rcpthosts
func RcpthostList() (hosts []core.RcptHost, err error) {
	return core.RcpthostGetAll()
}

// DKIM
func DkimEnable(domain string) (dkimConfig *core.DkimConfig, err error) {
	return core.DkimEnable(domain)
}
