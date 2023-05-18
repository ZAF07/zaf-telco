package app

import (
	"database/sql"
	"log"
	"net"

	httpserver "github.com/ZAF07/telco/api/http_server"
	"github.com/ZAF07/telco/config"
	"github.com/ZAF07/telco/datastore"
	"github.com/ZAF07/telco/internal/adapters/handlers"
	mockdb "github.com/ZAF07/telco/internal/adapters/repository/mock_db"
	postgresdb "github.com/ZAF07/telco/internal/adapters/repository/postgresDB"
	"github.com/ZAF07/telco/internal/core/ports"
	"github.com/ZAF07/telco/internal/core/services"
	"github.com/gin-contrib/cors"
)

const (
	POSTGRES = "postgres"
	TEST     = "test"
)

type App struct {
	HttpServer      ports.IHTTPServer
	RPCServer       ports.IRPCServer
	PrimaryDB       interface{}
	AppServices     TelcoServices
	AppHandlers     TelcoHandlers
	AppRepositories TelcoRepositories
	HTTPListener    net.Listener
	RPCListener     net.Listener
	AppConfig       *config.Config
}

type TelcoHandlers struct {
	TelcoAccountHanlder handlers.TelcoAccountHandler
	TelcoSettingHandler handlers.TelcoSettingHandler
}

type TelcoServices struct {
	TelcoAccountService ports.ITelcoAccountService
	TelcoSettingService ports.ITelcoSettingService
}

type TelcoRepositories struct {
	TelcoAccountRepo ports.ITelcoAccountRepository
	TelcoSettingRepo ports.ITelcoSettingRepository
}

func NewTelcoApplication(HTTPListener, RPCListener net.Listener, c *config.Config) *App {
	return &App{
		HTTPListener: HTTPListener,
		RPCListener:  RPCListener,
		AppConfig:    c,
	}
}

func (a *App) InitApplication() {

	a.initHTTPServer(a.HTTPListener, a.AppConfig.HTTPServerConfig, a)
	a.initPrimaryDB()
	a.initAppRepoAdapters()
	a.initAppServices()
	a.initHandlerAdapters()
	a.initRoutes()
}

func (a *App) initPrimaryDB() {
	switch a.AppConfig.Datastore.PrimaryDBType {
	case POSTGRES:
		conn := a.AppConfig.Datastore.GetPrimaryConnString()
		a.PrimaryDB = datastore.NewPostgresDB(conn)
	case TEST:
		log.Println("💡Using TEST DB")
		a.PrimaryDB = &sql.DB{}
	}

}

func (a *App) initAppRepoAdapters() {

	switch a.AppConfig.Datastore.PrimaryDBType {
	case POSTGRES:
		a.AppRepositories.TelcoAccountRepo = postgresdb.NewTelcoAccountRepoAdapter(a.PrimaryDB.(*sql.DB), "accounts")
		a.AppRepositories.TelcoSettingRepo = postgresdb.NewTelcoSettingRepoAdapter(a.PrimaryDB.(*sql.DB), "settings")
		return
	case TEST:
		log.Println("💡Using mock database adapters")
		a.AppRepositories.TelcoAccountRepo = mockdb.NewMockTelcoAccountRepoAdapter()
		a.AppRepositories.TelcoSettingRepo = mockdb.NewMockTelcoSettingRepoAdapter()
		return
	}
}

func (a *App) initAppServices() {
	a.AppServices.TelcoAccountService = services.NewTelcoAccountService(a.AppRepositories.TelcoAccountRepo)
	a.AppServices.TelcoSettingService = services.NewTelcoSettingsService(a.AppRepositories.TelcoSettingRepo)
}

func (a *App) initHandlerAdapters() {
	a.AppHandlers.TelcoAccountHanlder = *handlers.NewTelcoAccountHandler(a.AppServices.TelcoAccountService)
	a.AppHandlers.TelcoSettingHandler = *handlers.NewTelcoSettingHandler(a.AppServices.TelcoSettingService)
}

func (a *App) initHTTPServer(listener net.Listener, c *config.HTTPConfig, app *App) {
	k := httpserver.NewHTTPServer(a.HTTPListener, a.AppConfig)
	// g := ginserver.NewGinHTTPServer(listener, c.Addr, c.ReadTimeout, c.WriteTimeout)
	app.HttpServer = k
}

func (a *App) initRoutes() {
	router := a.HttpServer.(*httpserver.HTTPServer).HTTPHandler

	router.Use(cors.New(cors.Config{
		AllowCredentials: false,
		// AllowAllOrigins: true,
		// AllowOrigins: []string{"http://localhost:8080"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	telcoAccountHanlder := a.AppHandlers.TelcoAccountHanlder
	accounts := router.Group("accounts")
	{
		accounts.GET("/create-account", telcoAccountHanlder.CreateTelcoAccount)
	}

	telcoSettingHandler := a.AppHandlers.TelcoSettingHandler
	settings := router.Group("settings")
	{
		settings.GET("/", telcoSettingHandler.GetTelcoSetting)
	}

}

func (a *App) Start() {
	if a.HttpServer != nil {
		log.Println("💡 Starting HTTP SERVER")
		a.HttpServer.Start()
	}

	if a.RPCServer != nil {
		log.Println("💡 Starting RPC SERVER")
		a.RPCServer.Start()
	}
}
