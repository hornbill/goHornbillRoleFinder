package main

import "encoding/xml"

type xmlmcResponse struct {
	MethodResult string `xml:"status,attr"`
	Params       struct {
		SessionID string `xml:"sessionId"`
	} `xml:"params"`
	State stateStruct `xml:"state"`
}

//Roles - struct to hold all role data
type Roles struct {
	MethodResult string `xml:"status,attr"`
	Params       struct {
		RoleListItem []struct {
			Application string `xml:"application"`
			Description string `xml:"description"`
			Name        string `xml:"name"`
			PrivLevel   string `xml:"privLevel"`
			Type        string `xml:"type"`
			SysRights   RoleInfo
			AppRights   AppRights
		} `xml:"roleListItem"`
	} `xml:"params"`
	State stateStruct `xml:"state"`
}

//RoleInfo - struct to hold role description and system rights
type RoleInfo struct {
	MethodResult string `xml:"status,attr"`
	Params       struct {
		Application string `xml:"application"`
		Description string `xml:"description"`
		PrivLevel   string `xml:"privLevel"`
		RightA      struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightA"`
		RightB struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightB"`
		RightC struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightC"`
		RightD struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightD"`
		RightE struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightE"`
		RightF struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightF"`
		RightG struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightG"`
		RightH struct {
			Rights []Rights `xml:",any"`
		} `xml:"rightH"`
		Type string `xml:"type"`
	} `xml:"params"`
	State stateStruct `xml:"state"`
}

//AppRights - application specific rights
type AppRights struct {
	MethodResult string `xml:"status,attr"`
	Params       struct {
		Application []struct {
			Name   string `xml:"name"`
			RightA struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightA"`
			RightB struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightB"`
			RightC struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightC"`
			RightD struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightD"`
			RightE struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightE"`
			RightF struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightF"`
			RightG struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightG"`
			RightH struct {
				Rights []Rights `xml:",any"`
			} `xml:"rightH"`
		} `xml:"application"`
	} `xml:"params"`
	State stateStruct `xml:"state"`
}

//Rights - for dynamic unmarshalling of sys/app rights
type Rights struct {
	XMLName xml.Name `xml:""`
	Value   bool     `xml:",chardata"`
}

type stateStruct struct {
	Code     string `xml:"code"`
	ErrorRet string `xml:"error"`
}
