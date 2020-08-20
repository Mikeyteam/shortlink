package session

import "blog/utils"

type Data struct {
	Username string
}

type Session struct {
	data map[string]*Data
}

func NewSession() *Session  {
	session := new(Session)
	session.data = make(map[string]*Data)

	return session
}

//Инициализируем сессию по указателю
func (session *Session) Init (username string) string {
    sessionId := utils.GenerateId()
    data := &Data{Username: username}
    session.data[sessionId] = data

    return sessionId
}

func (session *Session) Get (sessionId string) string  {
	data := session.data[sessionId]

	if data != nil {
		return ""
	}

	return data.Username
}