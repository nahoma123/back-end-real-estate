package module

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/model"

	"github.com/golang-jwt/jwt/v4"
)

type AuthDetail struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

type UserModule interface {
	VerifyUserStatus(ctx context.Context, phone string) error
	VerifyToken(signingMethod jwt.SigningMethod, tokenString string) (bool, *jwt.MapClaims, error)
	GetUserStatus(ctx context.Context, Id string) (string, error)
	GetUser(ctx context.Context, Id string) (*model.User, error)
	Login(ctx context.Context, email, password string) (*AuthDetail, error)
	RegisterUser(ctx context.Context, profile *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, profile *model.User) (*model.User, error)
	GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error)
}

type EstateModule interface {
	GetFacultyServices(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyService, error)
	AddHouse(ctx context.Context, fService *model.House) (*model.House, error)
	CreateEstate(ctx context.Context, estate *model.Estate) (*model.Estate, error)
	CreateFacultyService(ctx context.Context, fService *model.FacultyService) (*model.FacultyService, error)
	AddEstateStaff(ctx context.Context, fService *model.Staff) (*model.Staff, error)
	AddHouseOwner(ctx context.Context, fService *model.HouseOwner) (*model.HouseOwner, error)
	GetHouses(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.House, error)
	GetEstateStaffs(ctx context.Context, filterPagination *constant.FilterPagination) ([]interface{}, error)
	AssignResident(ctx context.Context, fService *model.Resident) (*model.Resident, error)
	GetResidents(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Resident, error)
	GetHouseFees(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFee, error)
	CreateHouseFee(ctx context.Context, fService *model.HouseFee) (*model.HouseFee, error)

	GetHouseFeeSubscription(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFeeSubscription, error)
	AddHouseFeeSubscription(ctx context.Context, fService *model.HouseFeeSubscription) (*model.HouseFeeSubscription, error)

	GetFacultyServiceSubscriptions(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyServiceSubscription, error)
	AddFacultyServiceSubscription(ctx context.Context, fService *model.FacultyServiceSubscription) (*model.FacultyServiceSubscription, error)

	AddGuest(ctx context.Context, fService *model.Guest) (*model.Guest, error)
	GetGuests(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Guest, error)

	AddInvites(ctx context.Context, fService *model.Invite) (*model.Invite, error)
	GetInvites(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Invite, error)

	GetInviteByCode(ctx context.Context, inviteCode string) (*model.Invite, error)
	VerifyCode(ctx context.Context, inv string) (*model.Invite, error)

	UpdateInvite(ctx context.Context, invite *model.Invite) error
	FetchHouseFeeSubscription(ctx context.Context, month time.Time, estateId string) (*model.TenantReport, error)
	GetMonthlyReport(ctx context.Context, month time.Time, estateID string) (*model.MonthlyReport, error)
	FacultySubscriptionFeeReport(ctx context.Context, month time.Time, estateId string) ([]model.FacultyServiceMonthlyReport, error)
	InvitesReport(ctx context.Context, year int, estateId string) ([]model.InviteReport, error)

	LogBooking(ctx context.Context, code string, fService *model.BookingLog) (*model.BookingLog, error)

	UpdateEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error)
	AddEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error)
	GetEstateConfiguration(ctx context.Context, estateID string) (*model.EstateConfiguration, error)
	BanHouseOwnersForUnpaidServices(ctx context.Context) error
}

type GenericModule interface {
	UpdateOne(ctx constant.Context, id string) (interface{}, error)
	GetAny(cxt context.Context, colName string, model interface{}, filterPagination *constant.FilterPagination) ([]interface{}, error)
}
