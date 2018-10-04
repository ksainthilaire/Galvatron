package protocol 
import (
	"encoding/json"
)

type ProtocolRequired struct {
	CurrentVersion		int 	`json:"currentVersion"`
	RequiredVersion		int 	`json:"requiredVersion"`
}

type HelloConnectMessage struct {
	MessageType 		string
	Key					[]int		`json:"key"`
	Salt				string		`json:"salt"`
}

type AssetsVersionChecked struct {
	AssetsVersion		string	`json:"assetsVersion"`
	StaticDataVersion	string	`json:"staticDataVersion"`
}

type IdentificationSuccessMessage struct {
	AccountCreation 		int 	`json:"accountCreation"`
	AccountId 				int 	`json:"accountId"`
	CommunityId 			int 	`json:"communityId"`
	HasRights 				bool	`json:"hasRights"`
	Login 					string	`json:"login"`
	Nickname 				string	`json:"nickname"`
	SecretQuestion 			string 	`json:"secretQuestion"`
	SubscriptionEndDate 	int		`json:"subscriptionEndDate"`
	WasAlreadyConnected 	bool	`json:"wasAlreadyConnected"`
}

type GameServerInformations struct {
	CharactersCount 	int		`json:"charactersCount"`
	Completion 			int 	`json:"completion"`
	Date 				int 	`json:"date"`
	Id 					int 	`json:"id"`
	IsSelectable 		bool 	`json:"isSelectable"`
	Status 				int 	`json:"status"`
}

type ServersListMessage struct {
	Servers 	[]GameServerInformations `json:"gameServerInformations"`
}

type SelectedServerDataMessage struct {
	Address						string 	`json:"address"`
	CanCreateNewCharacter 		bool 	`json:"canCreateNewCharacter"`
	Port 						int		`json:"port"`
	ServerId 					int 	`json:"serverId"`
	Ticket 						string 	`json:"ticket"`
}

type CredentialsAcknowledgementMessage struct{}

type LoginQueueStatusMessage struct {
	Position 	int 	`json:"position"`
	Total 		int		`json:"total"`
}

func NewProtocolRequired(currentVersion int, requiredVersion int) byte[] {
	return json.Unmarshal(&ProtocolRequired{"ProtocolRequired", currentVersion, requiredVersion})
}

func NewIdenttificationSuccessMessage(
	accountCreation			int,
	accountId				int,
	communityId 			int,
	HasRights 				bool,
	Login 					string,
	Nickname				string,
	SecretQuestion 			string,
	SubscriptionEndDate		int,
	WasAlreadyConnected		bool
) byte[] {

	return json.Unmarshal(&IdentificationSuccessMessage(
		"IdentificationSuccessMessage",
		accountCreation,
		accountId,
		communityId,
		hasRights,
		login,
		nickname,
		secretQuestion,
		subscriptionEndDate,
		wasAlreadyConnected))
}

func NewAssetsVersionChecked(assetsVersion string, staticDataVersion string){
	return json.Unmarshal(&AssetsVersionChecked{"AssetsVersionChecked", assetsVersion, staticDataVersion})
}

func NewSelectedServerDataMessage(
	address 				string, 
	canCreateNewCharacter 	bool, 
	port 					int, 
	serverId				int, 
	ticket 					string
) byte[] {
	return json.Unmarshal(&SelectedServerDataMessage{
		"SelectedServerDataMessage", 
		address, 
		port, 
		serverId, 
		ticket})
}

func NewServersListMessage(servers []GameServerInformations) byte[] {
	return json.Unmarshal(&ServersListMessage{"ServersListMessage", servers})
}

func NewCredentialsAcknowledgementMessage() byte[] {
	return json.Unmarshal(&CredentialsAcknowledgementMessage{"CredentialsAcknowledgementMessage"})
}

func NewLoginQueueStatusMessage(position int, total int) byte[] {
	return json.Unmarshal(&LoginQueueStatusMessage{"LoginQueueStatusMessage", position, total})
}