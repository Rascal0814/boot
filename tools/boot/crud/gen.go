package crud

import (
	"embed"
	"fmt"
	"github.com/Rascal0814/boot/orm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed template/*.tpl
var tpl embed.FS

type GenCrud struct {
	db         *gorm.DB // db conn
	ModelPkg   string   // model package name
	OutputPath string   // generate file path
}

// splitScheme 将传入的DSN信息按照分隔符进行分开
func splitScheme(dsn string) (string, string, error) {
	segments := strings.SplitN(dsn, "://", 2)
	if len(segments) != 2 {
		return "", "", errors.New("the DSN mismatched the form driver://user:pass@host/database")
	}
	return segments[0], segments[1], nil
}

func NewGenCurd(dsn string, modelPkg string, output string) (*GenCrud, error) {
	dirver, s, err := splitScheme(dsn)
	if err != nil {
		return nil, err
	}
	connectDB, err := orm.ConnectDB(dirver, s)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(connectDB, &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return &GenCrud{db: db, ModelPkg: modelPkg, OutputPath: output}, nil
}

func (g *GenCrud) getTables() ([]*Table, error) {
	tables, err := g.db.Migrator().GetTables()
	if err != nil {
		return nil, err
	}
	var ts = make([]*Table, 0)
	// todo export naming to conf
	var ns = schema.NamingStrategy{}
	for _, t := range tables {
		if t == "schema_migrations" {
			continue
		}
		ts = append(ts, &Table{
			TableName: t,
			ModelName: ns.SchemaName(t),
			ModelPkg:  g.ModelPkg,
		})
	}
	return ts, nil
}

func (g *GenCrud) Gen() error {
	tables, err := g.getTables()
	if err != nil {
		return err
	}
	for _, t := range tables {
		modelPath := path.Join(g.OutputPath, fmt.Sprintf("%s.gen.go", t.TableName))
		tmp, err := template.ParseFS(tpl, "template/crud.tpl")
		if err != nil {
			return errors.Errorf("parse template file failed,%v", err)
		}
		if err := os.MkdirAll(filepath.Dir(modelPath), 0755); err != nil {
			return err
		}
		f, err := os.Create(modelPath)
		if err != nil {
			return errors.Errorf("created gen file failed,%v", err)
		}
		err = tmp.Execute(f, t)
		if err != nil {
			return errors.Errorf("generate go file failed,%v", err)
		}
	}

	return nil
}
