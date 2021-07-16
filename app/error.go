package app

import (
	"encoding/json"
	"fmt"
	"path"
	"reflect"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

const (
	Status             = "200"
	LocaleContext      = "locale"
	CodeSuccess        = "00000"
	DescriptionSuccess = "success"
)

type ErrorCode struct {
	Status      string            `json:"status"`
	MessageCode string            `json:"message_code"`
	Description LocaleDescription `json:"message_description"`
}

func (ec ErrorCode) Error() string {
	return ec.Description.EN
}

func (ec ErrorCode) WithLocale(c *gin.Context) ErrorCode {
	locale, ok := c.Value(LocaleContext).(string)
	if !ok {
		ec.Description.Locale = "th"
	}
	ec.Description.Locale = locale
	return ec
}

func (ec ErrorCode) WithFormat(a ...interface{}) ErrorCode {
	ec.Description.TH = fmt.Sprintf(ec.Description.TH, a...)
	ec.Description.EN = fmt.Sprintf(ec.Description.EN, a...)
	return ec
}

type LocaleDescription struct {
	TH     string
	EN     string
	Locale string
}

func (ld LocaleDescription) MarshalJSON() ([]byte, error) {
	if strings.ToLower(ld.Locale) == "th" {
		return json.Marshal(ld.TH)
	}
	return json.Marshal(ld.EN)
}

func (ld *LocaleDescription) UnmarshalJSON(data []byte) error {
	var res string
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	ld.EN = res
	ld.Locale = "en"
	return nil
}

type ErrorMessage struct {
	vn      *viper.Viper
	Configs Configs

	Success    ErrorCode
	BadRequest ErrorCode
	General    ErrorCode
	NoContent  ErrorCode
	NotFound   ErrorCode

	Invalid struct {
		Request     ErrorCode
		Transaction struct {
			DoesNotExist ErrorCode
			Amount       ErrorCode
			Format       ErrorCode
		}
	}

	Auth struct {
		//400
		InvalidUsernamePassword    ErrorCode
		InvalidCredential          ErrorCode
		InvalidRefreshToken        ErrorCode
		InvalidClientId            ErrorCode
		GetClaimsFail              ErrorCode
		InvalidClientSecret        ErrorCode
		InvalidClientType          ErrorCode
		AccessDenied               ErrorCode
		PermissionNotFound         ErrorCode
		LoginNotFound              ErrorCode
		SessionNotFound            ErrorCode
		SessionNotAvailable        ErrorCode
		InvalidResetPassword       ErrorCode
		InvalidNewPassword         ErrorCode
		RoleNotFound               ErrorCode
		InvalidOtp                 ErrorCode
		InvalidUsernameEmail       ErrorCode
		CompIntegrationNotFound    ErrorCode
		InvalidConfirmPassword     ErrorCode
		InvalidAcceptPrivacyNotice ErrorCode
		GetPrivacyNoticeBadRequest ErrorCode
		AuthorizationNotFound      ErrorCode
		AuthorizationExpiration    ErrorCode
		StatusInternalServerError  ErrorCode

		//401 internal
		TokenParseFail               ErrorCode
		InvalidToken                 ErrorCode
		TokenExpired                 ErrorCode
		TokenUsedBeforeAssign        ErrorCode
		InactiveLogin                ErrorCode
		InactiveComp                 ErrorCode
		GetPrivacyNoticeUnauthorized ErrorCode

		//404 notfound
		GetPrivacyNoticeNotFound ErrorCode

		//500
		GetLoginFail            ErrorCode
		DecodeHashPasswordFail  ErrorCode
		CreateAccessTokenFail   ErrorCode
		SigningAccessTokenFail  ErrorCode
		CreateRefreshTokenFail  ErrorCode
		SigningRefreshTokenFail ErrorCode
		AddNewSessionFail       ErrorCode
		RefreshingSessionFail   ErrorCode
		RevokeSessionFail       ErrorCode
		GetSessionFail          ErrorCode
		GetClientIdFail         ErrorCode
		GetClientSecretFail     ErrorCode
		GetSaltFail             ErrorCode
		ConvertTickFail         ErrorCode
		GetPrivateKeyFail       ErrorCode
		GetPublicKeyFail        ErrorCode
		ParsePrivateKeyFail     ErrorCode
		ParsePublicKeyFail      ErrorCode
		InitialClaimFail        ErrorCode
		SignedStringFail        ErrorCode
		HashingFail             ErrorCode
		InitialTokenFail        ErrorCode
		UpdateOtpFail           ErrorCode
		GetRoleFail             ErrorCode
		PasswordIsDuplicate     ErrorCode
		DecodeHashOtpFail       ErrorCode
		UpdateLoginHistoryFail  ErrorCode
		GetCompIntegrationFail  ErrorCode
		GetPrivacyNoticeFail    ErrorCode
		ParsePrivacyNoticeFail  ErrorCode
		AcceptPrivacyNoticeFail ErrorCode
		ConnectingLdapFail      ErrorCode
		BindingLdapFail         ErrorCode

		GetRolePermissionFail ErrorCode //w8 define
		GetCompanyTypeFail    ErrorCode
		GetCompanyFail        ErrorCode
	}

	Character struct {
		GetItemError             ErrorCode
		GetItemNotFound          ErrorCode
		GetCharacterInfoNotFound ErrorCode
		GetCharacterNotFound     ErrorCode
		NotEnoughBalance         ErrorCode
		GetBuffInfoNotFound      ErrorCode
		CharacterIDNotFound      ErrorCode
		GetItemListNotFound      ErrorCode
	}

	Farm struct {
		GetFarmNotFound                    ErrorCode
		GetPlantDexNotFound                ErrorCode
		GetItemNotFound                    ErrorCode
		UpdateItemPoolFail                 ErrorCode
		UpdateGrowUpFail                   ErrorCode
		GetInvInventoryNotFound            ErrorCode
		UpdateFarmFail                     ErrorCode
		UpdateInventoryFail                ErrorCode
		AddNewItemFail                     ErrorCode
		HarvestFailInvalidState            ErrorCode
		HarvestFailInvalidHarvestRemaining ErrorCode
		GetCharacterInfoNotFound           ErrorCode
		GetBuffNotFound                    ErrorCode
		UpdateQuantityFail                 ErrorCode
		DeleteInventToryFail               ErrorCode
		CheckRemainingNotFound             ErrorCode
		CheckRemainingHavePlant            ErrorCode
		CheckInventoryNotFound             ErrorCode
		CheckInventoryHaveNoSeed           ErrorCode
		CheckSeedNotFound                  ErrorCode
		CheckSeedOnlySeedPlant             ErrorCode
		GetPlantDexByItemIdNotFound        ErrorCode
		HarvestHornyFailInvalidState       ErrorCode

		ValidateFail struct {
			CheckFarm    ErrorCode
			CheckPlant   ErrorCode
			CheckWatered ErrorCode
		}
		UpdateFail struct {
			UpdateWatered ErrorCode
		}
	}

	Market struct {
		QuantityBadRequest      ErrorCode
		MarketIDBadRequest      ErrorCode
		GetItemError            ErrorCode
		GetItemNotFound         ErrorCode
		GetCharacterError       ErrorCode
		GetCharacterNotFound    ErrorCode
		LotteryNumberBadRequest ErrorCode
		GetLotteryError         ErrorCode
		LotteryIsBought         ErrorCode
		NotEnoughBalance        ErrorCode
		UpdateBalanceError      ErrorCode
		UpdateLotteryPoolError  ErrorCode
		InsertInventoryError    ErrorCode
		UpdateInventoryError    ErrorCode
		CommitError             ErrorCode
		
		ValidateFail           struct {
			ValidateQuantityFail	  ErrorCode
			ItemTypeCantSell          ErrorCode
			ItemQuantityCantSell      ErrorCode
			RefreshMandragoraBuffFail ErrorCode
			CheckBuffFail             ErrorCode
		}
		UpdateFail struct {
			ItemQuantityCantUpdate   ErrorCode
			ItemGoldCantUpdateC1     ErrorCode
			ItemGoldCantUpdateC2     ErrorCode
			ItemGoldCantUpdateNormal ErrorCode
			BuffCantUpdateC1         ErrorCode
			BuffCantUpdateC2         ErrorCode
			BuffCantUpdateRefresh    ErrorCode
		}
		InsertFail struct {
			InsertCharacter					ErrorCode
			InsertBuffFail 					ErrorCode
		}
		DeleteFail struct {
			DeleteInventoryFail 			ErrorCode
			DeleteBuffFail 					ErrorCode
		}
	}

	Register struct {
		ValidateFail struct {
			ValidateUsernameEmpty		ErrorCode
			ValidateUsernameLength		ErrorCode
			ValidatePasswordLength		ErrorCode
			ValidateEmailLength		    ErrorCode
			PasswordNotMatch  			ErrorCode
			DuplicateEmail    			ErrorCode
			DuplicateUsername 			ErrorCode
		}
		CreateFail struct {
			CreateUsername  ErrorCode
			CreateCharacter ErrorCode
			CreateInventory ErrorCode
			CreateFarm      ErrorCode
		}
	}

	Lottery struct{
		GetLotteryError ErrorCode
		GetGamuOwnerError ErrorCode
		GetLotteryCountError ErrorCode
		GetRandomLotteryError ErrorCode
		UpdateLotteryOwnerError ErrorCode
		InvalidLotteryInput ErrorCode
	}

	ConfigPath string
}

func (em *ErrorMessage) Init() error {
	vn := viper.New()
	vn.AddConfigPath(path.Join(em.Configs.ConfigPath, "general"))
	name := fmt.Sprintf("error.%s", em.Configs.Stage)
	log.Infof("c.Stage name: %s", name)
	vn.SetConfigName(name)

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	em.vn = vn
	em.mapping("", reflect.ValueOf(em).Elem())
	em.Success.Status = Status
	em.Success.MessageCode = CodeSuccess
	em.Success.Description.EN = DescriptionSuccess
	return nil
}

func (em ErrorMessage) mapping(name string, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		fi := v.Field(i)
		if fi.Kind() != reflect.Struct {
			continue
		}

		fn := underscore(v.Type().Field(i).Name)
		if name != "" {
			fn = fmt.Sprint(name, ".", fn)
		}

		if fi.Type().Name() == "ErrorCode" {
			fi.Set(reflect.ValueOf(em.ErrorCode(fn)))
			continue
		}
		em.mapping(fn, fi)
	}
}

func (em ErrorMessage) ErrorCode(name string) ErrorCode {
	rtn := ErrorCode{
		Status:      em.vn.GetString(fmt.Sprintf("%s.status", name)),
		MessageCode: em.vn.GetString(fmt.Sprintf("%s.code", name)),
		Description: LocaleDescription{
			TH: em.vn.GetString(fmt.Sprintf("%s.th", name)),
			EN: em.vn.GetString(fmt.Sprintf("%s.en", name)),
		},
	}
	return rtn
}

func (em ErrorMessage) InvalidRequestByName(name string) ErrorCode {
	return reflect.ValueOf(em.Invalid.Request).FieldByName(name).Interface().(ErrorCode)
}

func underscore(str string) string {
	runes := []rune(str)
	var out []rune
	for i := 0; i < len(runes); i++ {
		if i > 0 && (unicode.IsUpper(runes[i]) || unicode.IsNumber(runes[i])) && ((i+1 < len(runes) && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}
	return string(out)
}
