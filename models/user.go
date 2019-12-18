package models

import "gopkg.in/mgo.v2/bson"

// User object
type User struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `json:"name" bson:"name"`
	// FirstName   string        `bson:"firstName" json:"firstName"`
	// LastName    string        `bson:"lastName" json:"lastName"`
	// PhoneNumber string        `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	// Gender                       string        `bson:"gender,omitempty" json:"gender,omitempty"`
	EmailId string `bson:"emailId,omitempty" json:"emailId"`
	// Address                      `bson:",inline" json:",inline" mapstructure:",squash"`
	// Roles                        []bson.ObjectId `bson:"roles" json:"roles"`
	Password string `bson:"password" json:"-"`
	// UnEncPassword                string          `json:"password" bson:"-"`
	// ConfirmPassword              string          `json:"confirmPassword" bson:"-"`
	// PasswordToken                string          `json:"-" bson:"passwordToken"`
	// Pin                          string          `bson:"pin" json:"pin"`
	// DateOfBirth                  time.Time       `bson:"dateOfBirth,omitempty" json:"-"`
	// DOBStr                       string          `bson:"-" json:"dateOfBirth"`
	// DateOfJoining                time.Time       `bson:"dateOfJoining,omitempty" json:"-"`
	// DateOfJoiningStr             string          `bson:"-" json:"dateOfJoining"`
	// DateOfTermination            time.Time       `bson:"dateOfTermination,omitempty" json:"-"`
	// DateOfTerminationStr         string          `bson:"-" json:"dateOfTermination"`
	// TerminationReason            string          `bson:"terminationReason" json:"terminationReason"`
	// ImageData                    `bson:",inline" json:",inline"`
	// FirstScreen                  string `bson:"firstScreen" json:"firstScreen"`
	// MailStatus                   bool   `bson:"mailStatus" json:"mailStatus,omitempty"`
	// RegistrationStatus           bool   `bson:"registrationStatus" json:"registrationStatus,omitempty"` //deprecated
	// WebTheme                     string `bson:"webTheme,omitempty" json:"webTheme,omitempty"`
	// CreatedLog                   `bson:",inline" json:",inline"`
	// ActiveStatus                 bool            `bson:"activeStatus" json:"activeStatus"`
	// SyncWithInventory            bool            `json:"syncWithInventory" bson:"syncWithInventory"`
	// SyncImageWithInventory       bool            `json:"syncImageWithInventory" bson:"syncImageWithInventory"`
	// InventoryToken               string          `json:"inventoryLoginToken" bson:"inventoryLoginToken"`
	// Language                     bson.ObjectId   `json:"language,omitempty" bson:"language,omitempty"`
	// Payments                     []UserPayment   `json:"payments" bson:"payments"`
	// ClockInWhenPunchIn           bool            `json:"clockInWhenPunchIn" bson:"clockInWhenPunchIn"`
	// LoginToken                   string          `bson:"loginToken" json:"loginToken"`
	// LingaAPIToken                string          `bson:"lingaApiToken" json:"lingaApiToken"`
	// SuperAdmin                   bool            `json:"-" bson:"superAdmin"`
	// ShowTutorial                 bool            `json:"-" bson:"showTutorial"`
	// Stores                       []bson.ObjectId `json:"stores" bson:"stores"`
	// HasSignIn                    bool            `json:"hasSignIn" bson:"hasSignIn"`
	// LoginRooToken                string          `json:"-" bson:"loginRooToken"`
	// ClockInRequired              bool            `json:"clockInRequired" bson:"clockInRequired"`
	// Accounts                     []bson.ObjectId `json:"accounts" bson:"accounts"`
	// DoAutoCashierOut             bool            `json:"doAutoCashierOut" bson:"doAutoCashierOut"`
	// HasCashTip                   bool            `json:"hasCashTip" bson:"hasCashTip"`
	// CashTipPercentage            int             `json:"cashTipPercentage" bson:"cashTipPercentage"`
	// CashTipPercentageStr         string          `json:"cashTipPercentageStr" bson:"cashTipPercentageStr"`
	// EnterpriseReportNotification `json:"enterpriseReportNotification" bson:"enterpriseReportNotification"`
	// KioskUser                    bool `json:"kioskUser" bson:"kioskUser"`
}
