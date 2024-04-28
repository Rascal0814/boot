package orm

import (
	"fmt"
	"github.com/Rascal0814/boot/orm"
	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}

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

	db, err := gorm.Open(connectDB, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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
	for _, t := range tables {
		if t == "schema_migrations" {
			continue
		}
		ts = append(ts, &Table{
			TableName: t,
			ModelName: g.toSchemaName(t),
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
		if err := os.MkdirAll(filepath.Dir(modelPath), 0755); err != nil {
			return err
		}
		f, err := os.Create(modelPath)
		if err != nil {
			return errors.Errorf("created gen file failed,%v", err)
		}
		tmp, err := template.ParseFiles("./template/crud.tpl")
		if err != nil {
			return errors.Errorf("parse template file failed,%v", err)
		}
		err = tmp.Execute(f, t)
		if err != nil {
			return errors.Errorf("generate go file failed,%v", err)
		}
	}

	return nil
}

func (g *GenCrud) toSchemaName(name string) string {
	result := strings.ReplaceAll(cases.Title(language.English).String(strings.ReplaceAll(name, "_", " ")), " ", "")
	for _, initialism := range commonInitialisms {
		result = regexp.MustCompile(cases.Title(language.English).String(strings.ToLower(initialism))+"([A-Z]|$|_)").ReplaceAllString(result, initialism+"$1")
	}
	return result
}
