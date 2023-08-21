package repositories

import (
	"github.com/Mhakimamransyah/oauth/core/entities"
	"github.com/Mhakimamransyah/oauth/infrastructures/databases/mysql"
	"github.com/Mhakimamransyah/oauth/infrastructures/databases/mysql/tables"
	remoteurl "github.com/Mhakimamransyah/oauth/infrastructures/remote_url"
)

type OauthRepository struct {
	accessTokenRemoteUrl *remoteurl.AccessToken
	githubUserRemoteUrl  *remoteurl.GithubUserApi
	googleUserRemoteUrl  *remoteurl.GoogleUserApi
	mySqlDB              *mysql.Mysql
}

func (repo *OauthRepository) GetAccessToken(account int, code string) (string, error) {
	return repo.accessTokenRemoteUrl.FetchAccessToken(account, code)
}

func (repo *OauthRepository) GetUserInformation(account int, accessToken string) (*entities.User, error) {

	var user entities.User

	// github user information
	if account == entities.GithubAccount {
		if err := repo.githubUserRemoteUrl.FetchUserInformation(accessToken, &user); err != nil {
			return nil, err
		}
	}

	// google user information
	if account == entities.GoogleAccount {
		if err := repo.googleUserRemoteUrl.FetchUserInformation(accessToken, &user); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (repo *OauthRepository) StoreDetailUser(user *entities.User) error {

	var userTable = tables.User{}

	userTable.BindNewUser(user)
	userTable.CountLogin = 1

	db, cancel := repo.mySqlDB.GetDB()
	defer cancel()

	db.Begin()

	if err := db.Create(&userTable).Error; err != nil {
		db.Rollback()
		return err
	}

	db.Commit()

	return nil
}

func (repo *OauthRepository) FindUserByEmail(email string) (*entities.User, error) {

	var userTable tables.User

	db, cancel := repo.mySqlDB.GetDB()
	defer cancel()

	if err := db.Where("email = ?", email).Find(&userTable).Error; err != nil {
		return nil, err
	}

	return userTable.ConvertToUserEntity(), nil
}

func (repo *OauthRepository) IncrUserLoginCount(user *entities.User) error {

	var userTable tables.User

	userTable.BindNewUser(user)

	db, cancel := repo.mySqlDB.GetDB()
	defer cancel()

	if err := db.Model(&userTable).Where("email = ?", user.Email).Update("count_login", user.CountLogin+1).Error; err != nil {
		return err
	}

	return nil
}

func NewOauthRepository(accessToken *remoteurl.AccessToken, githubUser *remoteurl.GithubUserApi, googleUser *remoteurl.GoogleUserApi, mysql *mysql.Mysql) *OauthRepository {
	return &OauthRepository{
		accessTokenRemoteUrl: accessToken,
		githubUserRemoteUrl:  githubUser,
		googleUserRemoteUrl:  googleUser,
		mySqlDB:              mysql,
	}
}
