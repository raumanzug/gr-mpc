// Package services defines services infrastructure.
//
// Services are objects that exist during the entire runtime of the
// application.  Each service has to be enlisted in interfaces/context.go.
// Five services are listed in interfaces/context.go:
//
//   - Backend contains the service that takes care of controlling the
//     protocols;
//   - Lifecycle configures hooks used by smartphone operating systems
//     used by its energy management;
//   - Preferences takes care of the persistent settings of the application.
//     In the present application, this includes the connection data of all
//     the media rendering engines known to this app and the media rendering
//     server currently in use;
//   - ServerTabs manages server tabs in the server tabs UI;
//   - UI contains the components of the UI.
//
// For each service there is
//
//   - an interface contained in the interfaces directory, which lists all
//     the methods assigned to the service in question;
//
//   - a subdirectory contained in the service directory, which contains the
//     implementation of the service in question;
//
//   - the file service.go in this subdirectory contains a procedure
//     with the signature
//     New(context interfaces.IApplicationContext) interfaces.<service>,
//     which sets the relevant service to the default state.  The import list
//     must contain at least the following package:
//
//     import “mpc/interfaces”
//
//   - This procedure must be started in the file
//     services/application/context.go.
//     The struct context_t must contain an entry with an interface type
//     that relates to the service in question.  In the procedure
//     New() interfaces.IApplicationContext
//     this entry must be initialized by calling the function mentioned above
//     New(context IApplicationContext) interfaces.<service>.
//
//   - Moreover, some services have to be initialized.  Do this using
//     the procedure configuration.Configure
//
//     configuration.Configure(application.New())
//
//     for which we have to import
//
//     import "github.com/raumanzug/gr-mpc/configuration"
//
//     Cf. main procedure of this project.
//
// The application runs through two phases during its runtime:
//
// - an initialization phase,
//
// - an usage phase.
//
// In the first phase, the services are initialized.  The mentioned function
// New() interfaces.IApplicationContext is called and its return value is
// written to the variable ApplicationContext in the [mpc/globals] package.
// During this initialization phase, this interfaces.IApplicationContext
// is available everywhere because it is given to the service initialization
// routines New(context IApplicationContext) interfaces.<service> in
// the parameter context.  As soon as the application is initialized,
// you access it with the help of
//
//	globals.ApplicationContext.<service>().<method>(<params>...)
//
// to access the methods of the relevant service.  The package globals must
// have been imported:
//
//	import “mpc/globals”
package services
