package storage

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/model"
)

type DatabaseCollection string

// database collection constants
const (
	Users                      DatabaseCollection = "users"
	Estate                     DatabaseCollection = "estates"
	FacultyServiceSubscription DatabaseCollection = "faculty_services_subscriptions"
	Guest                      DatabaseCollection = "guests"
	Invite                     DatabaseCollection = "invites"
	FacultyService             DatabaseCollection = "faculty_services"
	Resident                   DatabaseCollection = "residents"
	HouseFee                   DatabaseCollection = "house_fees"
	HouseFeeSubscription       DatabaseCollection = "house_fee_subscriptions"
	Staff                      DatabaseCollection = "staffs"
	HouseOwner                 DatabaseCollection = "house_owners"
	House                      DatabaseCollection = "houses"
	BookingLog                 DatabaseCollection = "booking_logs"
	EstateConfiguration        DatabaseCollection = "estate_configurations"
	HouseOwnerInvite           DatabaseCollection = "house_owner_invites"
	Default                    DatabaseCollection = "defaults"
)

type UserStorage interface {
	Create(ctx context.Context, estate *model.User) (*model.User, error)
	Update(ctx context.Context, estate *model.User) (*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, string string) (*model.User, error)
	GetAll(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.User, error)
}

type EstateStorage interface {
	GetHouses(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.House, error)
	AddEstateStaff(ctx context.Context, fService *model.Staff) (*model.Staff, error)
	CreateFacultyService(ctx context.Context, fService *model.FacultyService) (*model.FacultyService, error)
	CreateEstate(ctx context.Context, estate *model.Estate) (*model.Estate, error)
	AddHouseOwner(ctx context.Context, fService *model.HouseOwner) (*model.HouseOwner, error)
	AddHouse(ctx context.Context, fService *model.House) (*model.House, error)
	GetFacultyServices(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyService, error)
	GetEstateStaffs(ctx context.Context, filterPagination *constant.FilterPagination) ([]interface{}, error)

	GetFacultyServiceSubscriptions(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyServiceSubscription, error)
	AddFacultyServiceSubscription(ctx context.Context, fService *model.FacultyServiceSubscription) (*model.FacultyServiceSubscription, error)
	FetchHouseFeeSubscription(ctx context.Context, month time.Time, estateId string) (*model.TenantReport, error)
	FacultySubscriptionFeeReport(ctx context.Context, month time.Time, estateId string) ([]model.FacultyServiceMonthlyReport, error)
	InvitesReport(ctx context.Context, year int, estateId string) ([]model.InviteReport, error)

	LogBooking(ctx context.Context, fService *model.BookingLog) (*model.BookingLog, error)
	UpdateEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error)
	AddEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error)
	GetEstateConfiguration(ctx context.Context, estateID string) (*model.EstateConfiguration, error)

	GetDefaults(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Default, error)
	AddDefault(ctx context.Context, dft *model.Default) (*model.Default, error)

	GetActiveDefault(ctx context.Context, residentID string) (*model.Default, error)
	IsResidentBanned(ctx context.Context, residentID string) (bool, error)
	BanHouseOwnersForUnpaidServices(ctx context.Context) error

	IsOwnerBanned(ctx context.Context, hoId string) (bool, error)
	GetActiveDefaultOwner(ctx context.Context, hoId string) (*model.Default, error)
}

type ResidentStorage interface {
	GetHouseFees(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFee, error)
	CreateHouseFee(ctx context.Context, fService *model.HouseFee) (*model.HouseFee, error)

	GetResidents(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Resident, error)
	GetResident(ctx context.Context, residentID string) (*model.Resident, error)
	UpdateResident(ctx context.Context, residentID string, rs *model.Resident) (*model.Resident, error)
	AssignResident(ctx context.Context, resident *model.Resident) (*model.Resident, error)

	GetHouseFeeSubscription(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFeeSubscription, error)
	AddHouseFeeSubscription(ctx context.Context, fService *model.HouseFeeSubscription) (*model.HouseFeeSubscription, error)

	AddGuest(ctx context.Context, fService *model.Guest) (*model.Guest, error)
	GetGuests(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Guest, error)

	AddInvites(ctx context.Context, fService *model.Invite) (*model.Invite, error)
	GetInvites(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Invite, error)

	GetInviteByCode(ctx context.Context, inviteCode string) (*model.Invite, error)
	UpdateInvite(ctx context.Context, invite *model.Invite) error
	GetMonthlyReport(ctx context.Context, month time.Time, estateId string) (*model.MonthlyReport, error)

	GetGuestDefaultsPerWeek(ctx context.Context, residentID string, weekDate time.Time) (int, error)
}

type GenericStorage interface {
	// UpdateOne(ctx constant.Context, id string) error
	DeleteOne(ctx context.Context, tableName string, field string, value interface{}) error
	GetOne(ctx context.Context, tableName string, data interface{}, field string, value interface{}) error
	CreateOne(ctx context.Context, tableName string, data interface{}) error
	UpdateOne(ctx context.Context, tableName string, updateData interface{}, field string, value interface{}) error
	GetAll(ctx context.Context, tableName string, data interface{}, filterPagination *constant.FilterPagination) (interface{}, error)
	// GetAny(cxt context.Context, colName string, model interface{}, filterPagination *constant.FilterPagination) ([]interface{}, error)
}
