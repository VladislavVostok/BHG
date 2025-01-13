package rpc

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/vmihailenco/msgpack/v5"
)

// Определения типов списка сессий Metasploit
type SessionListReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Token    string
}

type SessionListRes struct {
	ID          uint32 `msgpack:",omitempty"`
	Type        string `msgpack:"type"`
	TunnelLocal string `msgpack:"tunnel_local"`
	TunnelPeer  string `msgpack:"tunnel_peer"`
	ViaExploit  string `msgpack:"via_exploit"`
	ViaPayload  string `msgpack:"via_payload"`
	Description string `msgpack:"desc"`
	Info        string `msgpack:"info"`
	Workspace   string `msgpack:"workspace"`
	SessionHost string `msgpack:"session_host"`
	SessionPort int    `msgpack:"session_port"`
	Username    string `msgpack:"username"`
	UUID        string `msgpack:"uuid"`
	ExploitUUID string `msgpack:"exploit_uuid"`
}

// Определение типов входа и выхода в Metasploit
type loginReq struct {
	_msgpack struct{} `msgpack:",asArray"`
	Method   string
	Username string
	Password string
}

type loginRes struct {
	Result       string `msgpack:"result"`
	Token        string `msgpack:"token"`
	Error        bool   `msgpack:"error"`
	ErrorClass   string `msgpack:"error_class"`
	ErrorMessage string `msgpack:"error_message"`
}

type logoutReq struct {
	_msgpack    struct{} `msgpack:",asArray"`
	Method      string
	Token       string
	LogoutToken string
}

type logoutRes struct {
	Result string `msgpack:"result"`
}

// Определение клиента Metasploit
type Metasploit struct {
	host  string
	user  string
	pass  string
	token string
}

func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}

	if err := msf.Login(); err != nil {
		return nil, err
	}

	return msf, nil
}

// Обобщенный метод send() с повторно используемой сериализацией и десериализацией.
func (msf *Metasploit) send(req interface{}, res interface{}) error { // Метод send() получает параметры запроса и ответа типа interface().
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)                   //Далее для кодирования запроса применяется библиотека msgpack.
	dest := fmt.Sprintf("http://%s/api", msf.host)        //После кодирования на основе данных из Metasploit получателя, msf , создается целевой URL, явно указывая тип содержимого как binary/message-pack.
	r, err := http.Post(dest, "binary/message-pack", buf) // Устанавливаем для тела формат сериализованных данных.

	if err != nil {
		return err
	}

	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil { //Происходит декодирование тела ответа
		return err
	}
	return nil
}

// Реализация вызовов Metasploit API.
/* Метод Login() */
func (msf *Metasploit) Login() error {

	ctx := &loginReq{
		Method:   "auth.login",
		Username: msf.user,
		Password: msf.pass,
	}

	var res loginRes

	if err := msf.send(ctx, &res); err != nil {
		return err
	}

	msf.token = res.Token

	return nil
}

/* Метод Logout() */
func (msf *Metasploit) Logout() error {
	ctx := &logoutReq{
		Method:      "auth.logout",
		Token:       msf.token,
		LogoutToken: msf.token,
	}
	var res logoutRes
	if err := msf.send(ctx, &res); err != nil {
		return err
	}
	msf.token = ""
	return nil
}

func (msf *Metasploit) SessionList() (map[uint32]SessionListRes, error) {
	req := &SessionListReq{Method: "session.list", Token: msf.token}
	res := make(map[uint32]SessionListRes)
	if err := msf.send(req, &res); err != nil {
		return nil, err
	}
	for id, session := range res {
		session.ID = id
		res[id] = session
	}
	return res, nil
}
