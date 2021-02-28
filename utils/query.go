package utils

const (
	GET_USERS = `select u.id_user,u.no_ktp,u.Username,u.tanggal_lahir,p.id_pendidikan,p.pendidikan,u.id_pekerjaan,pk.pekerjaan,u.user_status
	FROM tb_users as u INNER JOIN tb_pendidikan as p
	ON u.id_pendidikan = p.id_pendidikan inner join
	tb_pekerjaan as pk on u.id_pekerjaan = pk.id_pekerjaan where u.user_status = 1`
	ADD_USER       = "INSERT INTO tb_users (id_user,username,no_ktp,tanggal_lahir,id_pekerjaan,id_pendidikan) values (?,?,?,?,?,?)"
	DELETE_USER    = "delete from tb_users where Id_user=?"
	GET_PROFESSION = `SELECT * FROM tb_pekerjaan`
	GET_EDUCATION  = `SELECT * FROM tb_pendidikan`
)
