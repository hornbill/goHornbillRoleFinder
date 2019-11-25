package main

import (
	"encoding/base64"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	apiLib "github.com/hornbill/goApiLib"
)

const (
	toolVer = "1.0.0"
)

var (
	right        string
	instanceName string
	userName     string
	userPassword string
	version      bool
	espXmlmc     *apiLib.XmlmcInstStruct
	roleData     Roles
)

func main() {
	flag.StringVar(&right, "right", "", "The right to search all roles for")
	flag.StringVar(&instanceName, "instance", "", "The instance name")
	flag.StringVar(&userName, "u", "", "The admin username")
	flag.StringVar(&userPassword, "p", "", "The admin password")
	flag.BoolVar(&version, "version", false, "Returns the version and ends")
	flag.Parse()
	//-- Used for Building
	if version {
		fmt.Printf("%v \n", toolVer)
		return
	}
	fmt.Println("===== Hornbill Role Search Utility v" + toolVer + " =====")

	//Try to login to the server if not succesfully exit
	success := login()
	if success != true {
		log.Fatal("Could not login to your Hornbill instance.")
	}
	defer logout()
	fmt.Println("Retrieving roles...")
	getRoles()

	if len(roleData.Params.RoleListItem) > 0 {
		//Get all role data in to maps
		fmt.Println("Retrieving all system and application rights for all roles...")
		for key, role := range roleData.Params.RoleListItem {
			roleData.Params.RoleListItem[key].SysRights = getRoleData(role.Name)
			roleData.Params.RoleListItem[key].AppRights = getRoleAppRights(role.Name)
		}

		fmt.Println("Searching for [" + right + "] in all system and application rights...")
		//Now search all roles for specified right
		for _, role := range roleData.Params.RoleListItem {

			//System Rights
			//RightA
			for _, sysRight := range role.SysRights.Params.RightA.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "a." + rightName + "]")
				}
			}
			//RightB
			for _, sysRight := range role.SysRights.Params.RightB.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "b." + rightName + "]")
				}
			}
			//RightC
			for _, sysRight := range role.SysRights.Params.RightC.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "c." + rightName + "]")
				}
			}
			//RightD
			for _, sysRight := range role.SysRights.Params.RightD.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "d." + rightName + "]")
				}
			}
			//RightE
			for _, sysRight := range role.SysRights.Params.RightE.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "e." + rightName + "]")
				}
			}
			//RightF
			for _, sysRight := range role.SysRights.Params.RightF.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "f." + rightName + "]")
				}
			}
			//RightG
			for _, sysRight := range role.SysRights.Params.RightG.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "g." + rightName + "]")
				}
			}
			//RightH
			for _, sysRight := range role.SysRights.Params.RightH.Rights {
				rightName := sysRight.XMLName.Local
				if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
					appStr := ""
					if role.Application != "" {
						appStr = " in Application [" + role.Application + "]"
					}
					color.HiBlue("Role [" + role.Name + "]" + appStr + " contains System Right [" + "h." + rightName + "]")
				}
			}

			//App Rights
			//RightA
			for _, appRight := range role.AppRights.Params.Application {
				for _, sysRight := range appRight.RightA.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "a." + rightName + "]")
					}
				}
				//RightB
				for _, sysRight := range appRight.RightB.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "b." + rightName + "]")
					}
				}
				//RightC
				for _, sysRight := range appRight.RightC.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "c." + rightName + "]")
					}
				}
				//RightD
				for _, sysRight := range appRight.RightD.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "d." + rightName + "]")
					}
				}
				//RightE
				for _, sysRight := range appRight.RightE.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "e." + rightName + "]")
					}
				}
				//RightF
				for _, sysRight := range appRight.RightF.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "f." + rightName + "]")
					}
				}
				//RightG
				for _, sysRight := range appRight.RightG.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "g." + rightName + "]")
					}
				}
				//RightH
				for _, sysRight := range appRight.RightH.Rights {
					rightName := sysRight.XMLName.Local
					if strings.Contains(strings.ToLower(rightName), strings.ToLower(right)) && sysRight.Value == true {
						appStr := ""
						if role.Application != "" {
							appStr = " in Application [" + role.Application + "]"
						}
						color.Green("Role [" + role.Name + "]" + appStr + " contains App Right [" + "h." + rightName + "]")
					}
				}
			}
		}
	} else {
		color.Red("No Roles found!")
	}
}

func getRoles() {
	browseRoles, err := espXmlmc.Invoke("admin", "roleGetList")
	if err != nil {
		color.Red("Call to roleGetList failed:", fmt.Sprintf("%v", err))
		os.Exit(0)
	}
	var xmlRespon Roles
	err = xml.Unmarshal([]byte(browseRoles), &xmlRespon)
	if err != nil {
		color.Red("Unmarshalling of data for roleGetList failed:", fmt.Sprintf("%v", err))
	}
	if xmlRespon.MethodResult != "ok" {
		color.Red("roleGetList was unsuccessful: "+xmlRespon.State.ErrorRet, "error")
	}
	roleData = xmlRespon
}

func getRoleData(role string) RoleInfo {
	espXmlmc.SetParam("role", role)
	getRole, err := espXmlmc.Invoke("admin", "roleGetInfo")
	if err != nil {
		color.Red("Call to roleGetInfo failed:", fmt.Sprintf("%v", err))
		os.Exit(0)
	}
	var xmlRespon RoleInfo
	err = xml.Unmarshal([]byte(getRole), &xmlRespon)
	if err != nil {
		color.Red("Unmarshalling of data for roleGetInfo failed:", fmt.Sprintf("%v", err))
	}
	if xmlRespon.MethodResult != "ok" {
		color.Red("roleGetInfo was unsuccessful: "+xmlRespon.State.ErrorRet, "error")
	}
	return xmlRespon
}

func getRoleAppRights(role string) AppRights {
	espXmlmc.SetParam("role", role)
	getRole, err := espXmlmc.Invoke("admin", "roleGetApplicationList")
	if err != nil {
		color.Red("Call to roleGetApplicationList failed:", fmt.Sprintf("%v", err))
		os.Exit(0)
	}
	var xmlRespon AppRights
	err = xml.Unmarshal([]byte(getRole), &xmlRespon)
	if err != nil {
		color.Red("Unmarshalling of data for roleGetApplicationList failed:", fmt.Sprintf("%v", err))
	}
	if xmlRespon.MethodResult != "ok" {
		color.Red("roleGetApplicationList was unsuccessful: "+xmlRespon.State.ErrorRet, "error")
	}
	return xmlRespon
}

//login - Starts a new ESP session
func login() bool {

	espXmlmc = apiLib.NewXmlmcInstance(instanceName)
	espXmlmc.SetParam("userId", userName)
	espXmlmc.SetParam("password", base64.StdEncoding.EncodeToString([]byte(userPassword)))
	XMLLogin, err := espXmlmc.Invoke("session", "userLogon")
	if err != nil {
		color.Red("Error returned when attempting to run Login API call.")
		fmt.Println(err)
		return false
	}
	var xmlRespon xmlmcResponse
	err = xml.Unmarshal([]byte(XMLLogin), &xmlRespon)
	if err != nil {
		color.Red("Error returned when attempting to unmarshal Login API call response.")
		fmt.Println(err)
		return false
	}
	if xmlRespon.MethodResult != "ok" {
		color.Red("Error returned when attempting to log in to your instance: " + xmlRespon.State.ErrorRet)
		fmt.Println(xmlRespon)
		return false
	}
	return true
}

//logout - Log out of ESP
func logout() {
	espXmlmc.Invoke("session", "userLogoff")
}
