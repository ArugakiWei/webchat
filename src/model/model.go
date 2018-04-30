package model


type Message struct {

	Id        	int       	`json:"-" orm:"column(id);pk"`
	Body		string		`json:"body" orm:"column(body)"`
	UserName	string		`json:"username" orm:"column(username)"`
	Time		string		`json:"time" orm:"column(time)"`
}

func NewMessage(userName, body, time string)  *Message{

	return &Message{
		Body:		body,
		UserName:	userName,
		Time:		time,
	}
}

type MessageResponse struct {

	Response	[]Message	`json:"messages"`
}

func NewMessageResponse(msg []Message)  *MessageResponse{

	return &MessageResponse{
		Response:	msg,
	}
}

type User struct {

	Id        	int       	`json:"-" orm:"column(id);pk"`
	UserName	string		`json:"username" orm:"column(username)"`
	PassWord	string		`json:"password" orm:"column(password)"`
	IsOnline	string		`json:"isonline" orm:"column(isonline)"`

}

func NewUserResponse(user []UserRes)  *UserResponse{

	return &UserResponse{
		Response:	user,
	}
}

type UserRes struct {

	UserName	string		`json:"username"`
	IsOnline	string		`json:"isonline"`
}

type UserResponse struct {

	Response	[]UserRes		`json:"users"`
}

type File struct {

	Id        	int       	`json:"-" orm:"column(id);pk"`
	FileName 	string		`json:"filename" orm:"column(filename)"`
	FilePath	string		`json:"filepath" orm:"column(filepath)"`
}

type FileResponse struct {

	Response	[]File		`json:"files"`
}

func NewFileResponse(file []File)  *FileResponse{

	return &FileResponse{
		Response:	file,
	}
}

type InfoResponse struct {

	Info	string		`json:"info"`
}

func NewInfoResponse(info string)  *InfoResponse{

	return &InfoResponse{
		Info:	info,
	}
}
