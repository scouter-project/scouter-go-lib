package configure

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Configure struct {
	confMap          map[string]string
	configFile       string
	filesize         int64
	fileModifiedtime time.Time
}

var configOnce sync.Once
var configure *Configure

func NewConfigure(file string) *Configure {
	configOnce.Do(func() {
		configure = &Configure{}
		configure.configFile = file
		configure.filesize = 0
		configure.confMap = make(map[string]string)
		configure.LoadEnv()
		configure.loadFile()
		configure.Start()
	})
	return configure
}

func GetConfigure() *Configure {
	if configure == nil {
		fmt.Println("configure is nil. Before GetConfigure, call NewConfigure() with config file path)")
		os.Exit(0)
	}
	return configure
}

var stopRunning = make(chan bool)

// Start is a method for reading configuraton.
func (c *Configure) Start() {
	initialStat, err := os.Stat(c.configFile)
	if err == nil {
		c.filesize = initialStat.Size()
		c.fileModifiedtime = time.Now()
	}
	go func() {
		for {
			select {
			case <-stopRunning:
				break
			default:
				c.loadConfig()

			}
			time.Sleep(10 * time.Second)
		}

	}()
}

// Stop is a method for stopping read configuration.
func Stop() {
	stopRunning <- true

}

func (c *Configure) LoadEnv() {
	for _, v := range os.Environ() {
		pair := strings.Split(v, "=")
		c.confMap[pair[0]] = pair[1]
	}
}

func (c *Configure) Put(key string, value string) {
	c.confMap[key] = value
}

func (c *Configure) Get(key string) string {
	if value, ok := c.confMap[key]; ok {
		return value
	}
	return ""

}
func (c *Configure) loadFile() (err error) {
	file, err := os.OpenFile(c.configFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if pos := strings.Index(line, "="); pos >= 0 {
			key := strings.TrimSpace(line[:pos])
			// 주석처리된 라인은 config에서 제외
			if len(key) > 0 && !strings.HasPrefix(key, "#") {
				value := ""
				if len(line) > pos {
					value = line[pos+1:]
				}
				c.Put(key, value)
			}
		}
	}
	return nil
}

// conf 파일에 파라미터 추가
func (c *Configure) Append(key string, value string) error {
	if _, ok := c.confMap[key]; !ok {
		file, err := os.OpenFile(c.configFile, os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString("\n" + key + "=" + value)
		file.Sync()
		c.Put(key, value)
	}
	return nil
}

// ReadStringValue returns string type config value
func (c *Configure) ReadStringValue(key string, def string) string {
	value := c.Get(key)
	if value == "" {
		return def
	}
	return value
}

// ReadIntValue returns int type config value
func (c *Configure) ReadIntValue(key string, def int) int {
	v := c.Get(key)
	if v == "" {
		return def
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return value
}

// ReadBoolValue returns bool type config value
func (c *Configure) ReadBoolValue(key string, def bool) bool {
	v := c.Get(key)
	if v == "" {
		return def
	}
	value, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return value

}

func (c *Configure) loadConfig() {
	if c.checkFile(c.configFile) {
		c.loadFile()
	}

}

func (c *Configure) checkFile(filePath string) bool {
	stat, err := os.Stat(c.configFile)
	if err != nil {
		return false
	}

	if stat.Size() != c.filesize || stat.ModTime() != c.fileModifiedtime {
		c.filesize = stat.Size()
		c.fileModifiedtime = stat.ModTime()
		return true
	}
	return false
}

func (c *Configure) ToString() string {
	s := ""
	for k, v := range c.confMap {
		s += k + ":" + v + "\n"

	}
	return s
}
