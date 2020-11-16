package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/BillotP/gorenty/v2/service"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

const (
	// DBHOST is the database host
	DBHOST = "192.168.1.34"
	// DBPORT is the database port
	DBPORT = "8529"
	// DBPROTOCOL is the database url protocol
	DBPROTOCOL = "https"
	// DBAPIUSER is the database user name
	DBAPIUSER = "michel"
	// DBPAPIASS is the database user password
	DBPAPIASS = "D8bFJaufwGdUCSC"
	// DBUSER is the database user name
	DBUSER = "root"
	// DBPASS is the database user password
	DBPASS = ""
	// DBNAME is the database name
	DBNAME = "renty-dev"
)

// Edge is db graph edge
type Edge struct {
	From string `json:"_from"`
	To   string `json:"_to"`
}

func seedDatabasesCollectionsAndIndexes(ctx context.Context, c driver.Client) error {
	for dbname, collections := range service.AllCollections {
		db, err := c.CreateDatabase(ctx, dbname, nil)
		if err != nil && !driver.IsConflict(err) {
			return err
		} else if driver.IsConflict(err) {
			db, err = c.Database(ctx, DBNAME)
			if err != nil {
				return err
			}
		}
		for _, el := range collections {
			col, err := db.CreateCollection(ctx, el, nil)
			if err != nil && !driver.IsConflict(err) {
				return err
			} else if driver.IsConflict(err) {
				col, err = db.Collection(ctx, el)
				if err != nil {
					return err
				}
			}
			fmt.Printf("Created collection %s.%s\n", dbname, el)
			switch el {
			case "locations":
				var idx driver.Index
				if idx, _, err = col.EnsureGeoIndex(ctx, []string{
					"geometry",
				}, &driver.EnsureGeoIndexOptions{
					GeoJSON: true,
				}); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create geo index %s on %s.'geometry' field\n", idx.Name(), el)
				break
			case "labels":
				var idx driver.Index
				if idx, _, err = col.EnsureFullTextIndex(ctx, []string{
					"value",
				}, nil); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create full text index %s on %s.'value' field\n", idx.Name(), el)
				break
			case "users":
				var idx driver.Index
				if idx, _, err = col.EnsurePersistentIndex(ctx, []string{
					"pseudo",
				}, &driver.EnsurePersistentIndexOptions{
					Unique: true,
				}); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create persistent unique index %s on %s.'pseudo' field\n", idx.Name(), el)
				break
			case "emails":
				var idx driver.Index
				if idx, _, err = col.EnsurePersistentIndex(ctx, []string{
					"value",
				}, &driver.EnsurePersistentIndexOptions{
					Unique: true,
				}); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create persistent unique index %s on %s.'value' field\n", idx.Name(), el)
				break
			case "phones":
				var idx driver.Index
				if idx, _, err = col.EnsurePersistentIndex(ctx, []string{
					"value",
				}, &driver.EnsurePersistentIndexOptions{
					Unique: true,
				}); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create persistent unique index %s on %s.'value' field\n", idx.Name(), el)
				break
			case "email_templates":
				var idx driver.Index
				if idx, _, err = col.EnsurePersistentIndex(ctx, []string{
					"label",
				}, &driver.EnsurePersistentIndexOptions{
					Unique: true,
				}); err != nil && !driver.IsConflict(err) {
					return err
				}
				fmt.Printf("Create persistent unique index %s on %s.'label' field\n", idx.Name(), el)
				break
			}
		}
		if dbname == "renty-dev" {
			allfields := true
			// asc := driver.ArangoSearchSortDirectionAsc
			if _, err = db.CreateArangoSearchView(ctx, "rentoffersearch", &driver.ArangoSearchViewProperties{
				Links: driver.ArangoSearchLinks{
					"labels": driver.ArangoSearchElementProperties{
						Analyzers:        []string{"identity"},
						IncludeAllFields: &allfields,
						Fields: driver.ArangoSearchFields{
							"value": driver.ArangoSearchElementProperties{
								Analyzers: []string{"text_en"},
							},
						},
					},
					"rentoffers": driver.ArangoSearchElementProperties{
						Analyzers:        []string{"identity"},
						IncludeAllFields: &allfields,
					},
				},
			}); err != nil && !driver.IsConflict(err) {
				return err
			}
		}
	}

	return nil
}

func seedUserAndRights(ctx context.Context, c driver.Client) error {
	var err error
	var usr driver.User
	if usr, err = c.CreateUser(ctx, DBAPIUSER, &driver.UserOptions{
		Password: DBPAPIASS,
	}); err != nil && !driver.IsConflict(err) {
		return err
	} else if driver.IsConflict(err) {
		if usr, err = c.User(ctx, DBAPIUSER); err != nil {
			return err
		}
	}
	fmt.Printf("Created db user %s\n", DBAPIUSER)
	for _, dbname := range service.AllDatabases {
		var db driver.Database
		if db, err = c.Database(ctx, dbname); err != nil {
			return err
		}
		if err = usr.SetDatabaseAccess(ctx, db, driver.GrantReadWrite); err != nil {
			return err
		}
		fmt.Printf("Gived user %s readwrite rights on %s db\n", DBAPIUSER, db.Name())
	}

	return nil
}

func main() {
	var err error
	var c driver.Client
	var db driver.Database
	var conn driver.Connection
	ctx := context.Background()
	endpoint := fmt.Sprintf("%s://%s:%s", DBPROTOCOL, DBHOST, DBPORT)
	if conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{endpoint},
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}); err != nil {
		log.Fatal(err)
	}
	if c, err = driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUSER, DBPASS),
	}); err != nil {
		log.Fatal(err)
	}
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "mock":
			if db, err = c.Database(ctx, DBNAME); err != nil {
				log.Fatal(err)
			}
			if err := seedDatas(ctx, db); err != nil {
				log.Fatal(err)
			}
			if err := seedEdges(ctx, db); err != nil {
				log.Fatal(err)
			}
			return
		case "drop":
			for _, dbname := range service.AllDatabases {
				var db driver.Database
				if db, err = c.Database(ctx, dbname); err != nil {
					log.Fatal(err)
				}
				if err = db.Remove(ctx); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Database %s successfully dropped\n", DBNAME)
			}
			var usr driver.User
			if usr, err = c.User(ctx, DBAPIUSER); err != nil {
				log.Fatal(err)
			}
			if err = usr.Remove(ctx); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("DB User %s successfully removed\n", DBAPIUSER)
			return
		default:
			log.Fatalf("%s: unknow command", os.Args[1])
		}
	}
	if err := seedDatabasesCollectionsAndIndexes(ctx, c); err != nil {
		log.Fatal(err)
	}
	if err := seedUserAndRights(ctx, c); err != nil {
		log.Fatal(err)
	}
}
