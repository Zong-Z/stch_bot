package betypes

import "testing"

func TestUser_IsSuitableAge(t *testing.T) {
	u := User{Age: "1", InterlocutorAge: "2"}
	u2 := User{Age: "2", InterlocutorAge: "1"}
	if !u.IsSuitableAge(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: "1", InterlocutorAge: "2"}
	u2 = User{Age: "1", InterlocutorAge: "2"}
	if u.IsSuitableAge(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: "2", InterlocutorAge: "2"}
	u2 = User{Age: "1", InterlocutorAge: "2"}
	if u.IsSuitableAge(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: "2", InterlocutorAge: UserNil}
	u2 = User{Age: UserNil, InterlocutorAge: UserNil}
	if !u.IsSuitableAge(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: UserNil, InterlocutorAge: "2"}
	u2 = User{Age: "2", InterlocutorAge: "1"}
	if u.IsSuitableAge(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: "2", InterlocutorAge: UserNil}
	u2 = User{Age: "1", InterlocutorAge: "2"}
	if !u.IsSuitableAge(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: UserNil, InterlocutorAge: UserNil}
	u2 = User{Age: "1", InterlocutorAge: "2"}
	if u.IsSuitableAge(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{Age: UserNil, InterlocutorAge: UserNil}
	u2 = User{Age: "1", InterlocutorAge: "2"}
	if u.IsSuitableAge(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}
}

func TestUser_IsSuitableCity(t *testing.T) {
	u := User{City: "1", InterlocutorCity: "2"}
	u2 := User{City: "2", InterlocutorCity: "1"}
	if !u.IsSuitableCity(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: "1", InterlocutorCity: "2"}
	u2 = User{City: "1", InterlocutorCity: "2"}
	if u.IsSuitableCity(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: "2", InterlocutorCity: "2"}
	u2 = User{City: "1", InterlocutorCity: "2"}
	if u.IsSuitableCity(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: "2", InterlocutorCity: UserNil}
	u2 = User{City: UserNil, InterlocutorCity: UserNil}
	if !u.IsSuitableCity(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: UserNil, InterlocutorCity: "2"}
	u2 = User{City: "2", InterlocutorCity: "1"}
	if u.IsSuitableCity(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: "2", InterlocutorCity: UserNil}
	u2 = User{City: "1", InterlocutorCity: "2"}
	if !u.IsSuitableCity(u2) {
		t.Errorf("SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: UserNil, InterlocutorCity: UserNil}
	u2 = User{City: "1", InterlocutorCity: "2"}
	if u.IsSuitableCity(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}

	u = User{City: UserNil, InterlocutorCity: UserNil}
	u2 = User{City: "1", InterlocutorCity: "2"}
	if u.IsSuitableCity(u2) {
		t.Errorf("NOT SUITABLE CASE. User 1: %v. User 2: %v.", u, u2)
	}
}
