package linkv_sdk

import (
	"github.com/linkv-io/go-sdk/im"
)

type IMGroupInfo = im.GroupInfo
type IMGroup = im.Group
type IMGroupUser = im.GroupUser

type IMGroupStatus = im.GroupStatus
type IMGroupUserStatus = im.GroupUserStatus

var (
	IMStatusGroupUnBlock     = im.GroupUnBlock
	IMStatusGroupBlock       = im.GroupBlock
	IMStatusGroupUserDisable = im.GroupUserDisable
	IMStatusGroupUserEnable  = im.GroupUserEnable
)

type LvIM interface {
	GetTokenByThirdUID(thirdUID string) (imToken string, err error)
	PushConverseData(fromUID, toUID, content, objectName, toAppID, toUserExtSysUserID string, args ...string) (requestID string, err error)
	PushConverseDatas(fromUID string, toUIDs []string, content, objectName string, args ...string) (requestID string, failUIDs []string, err error)
	SendRoomMessage(fromUID, toRoomID, content, objectName string) (requestID string, err error)
	SendEventMsg(fromUID, toUID, content, objectName string) (requestID string, err error)
	UserBlock(userIDs []string, min int64) (err error)
	UserStatus(userID string) (bOnline bool, requestID string, err error)
	AddUserBlack(fromUID string, userIDs []string) (requestID string, failUIDs []string, err error)
	RemoveUserBlack(fromUID string, userIDs []string) (requestID string, failUIDs []string, err error)

	GroupInfo(groupID string) (*IMGroupInfo, string, error)
	GroupModifyName(adminUserID, groupID, groupName string) (string, error)
	GroupModifyTopic(adminUserID, groupID, topic string) (string, error)
	GroupModifyAdmin(adminUserID, toUserID, groupID string) (string, error)

	GroupDissolve(adminUserID, groupID string, status IMGroupStatus) (string, error)
	GroupKickUser(adminUserID, kickUserID, groupID string) (string, error)
	GroupGagUser(adminUserID, gagUserID, groupID string, status IMGroupUserStatus) (string, error)
	GroupGagAllUser(adminUserID, groupID string, status IMGroupUserStatus) (string, error)

	GroupCreate(userID, nickName, groupID, groupName string) (string, error)
	GroupJoin(userID, nickName, groupID string) (string, error)
	GroupExit(userID, groupID string) (string, error)
	GroupList(userID string, page, size int) ([]IMGroup, string, error)
	GroupSetNotifiable(userID, groupID string, status IMGroupUserStatus) (string, error)

	GroupUserList(groupID string, page, size int) ([]IMGroupUser, string, error)
}

func NewLvIM(operatorID string) LvIM {
	return im.New(operatorID)
}
