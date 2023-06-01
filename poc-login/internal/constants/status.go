package constants

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type UserStatus string

const (
	UserStatusActive   = UserStatus("ACTIVE")
	UserStatusInactive = UserStatus("INACTIVE")
)

func (u UserStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UserStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = UserStatus(cloneBytes(source))
	return nil
}

func cloneBytes(source []byte) []byte {
	dst := make([]byte, len(source))
	copy(dst, source)
	return dst
}

type PermissionStatus string

const (
	PermissionStatusActive   = PermissionStatus("ACTIVE")
	PermissionStatusInactive = PermissionStatus("INACTIVE")
)

func (u PermissionStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *PermissionStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = PermissionStatus(cloneBytes(source))
	return nil
}

type UserAccessStatus string

const (
	UserAccessStatusActive   = UserAccessStatus("ACTIVE")
	UserAccessStatusInactive = UserAccessStatus("INACTIVE")
)

func (u UserAccessStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *UserAccessStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = UserAccessStatus(cloneBytes(source))
	return nil
}

type RoleStatus string

const (
	RoleStatusActive   = RoleStatus("ACTIVE")
	RoleStatusInactive = RoleStatus("INACTIVE")
)

func (u RoleStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *RoleStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = RoleStatus(cloneBytes(source))
	return nil
}

type RolePermissionStatus string

const (
	RolePermissionStatusActive   = RolePermissionStatus("ACTIVE")
	RolePermissionStatusInactive = RolePermissionStatus("INACTIVE")
)

func (u RolePermissionStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *RolePermissionStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = RolePermissionStatus(cloneBytes(source))
	return nil
}

type RefreshTokenStatus string

const (
	RefreshTokenStatusActive   = RefreshTokenStatus("ACTIVE")
	RefreshTokenStatusInactive = RefreshTokenStatus("EXPIRED")
)

func (u RefreshTokenStatus) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *RefreshTokenStatus) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	*u = RefreshTokenStatus(cloneBytes(source))
	return nil
}
