package dataprofile

type Profile struct {
	Nama   string
	Nim    string
	Alamat string
}

// var profile1 = Profile{"Hanley", "2502010566", "Argo Tunggal 19 Lawang"}
// var profile2 = Profile{"Yunanda", "2401056602", "Araya Mansion 8 Pakis"}
// var profile3 = Profile{"Gusti", "1234", "Malang"}
// var profile4 = Profile{"Pangestu", "321", "Kabupaten Malang"}

var listProfile = [4]Profile{
	Profile{"Hanley", "2502010566", "Argo Tunggal 19 Lawang"},
	Profile{"Yunanda", "2401056602", "Araya Mansion 8 Pakis"},
	Profile{"Gusti", "1234", "Malang"},
	Profile{"Pangestu", "321", "Kabupaten Malang"},
}

func GetDataProfile(nama string) Profile {
	idx := 0
	for i := 0; i < 4; i++ {
		if nama == listProfile[i].Nama {
			idx = i
		}
	}
	return listProfile[idx]
}

func GetAllDataProfile() [4]Profile {
	return listProfile
}
