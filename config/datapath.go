package config

const (
	Scripts_path string = "/scripts2run/"
	Binaries_path string = "/bin2copy/"
	//Binaries_baseurl string = "https://github.com/kmahyyg/ki11-aegis/releases/download/"
	//Binaries_tag string = "v1.0.0/"
	Iptables_cmdprefix string = "iptables -w -A INPUT -s "
	Iptables_cmdsuffix string = " -j DROP"
	Apt_PreStart = "apt update -y "
	Apt_Inst = "apt install -y --no-install-recommends "
	Postclean_base string = "rm -rf "
	Extract_base = "tar -x -J -C /usr/local/bin -f "
)

//var Must_preClean_Scripts = []string{
//	"http://update.aegis.aliyun.com/download/uninstall.sh",
//	"http://update.aegis.aliyun.com/download/quartz_uninstall.sh",
//}

var Must_postClean_Scripts = []string{
	"dpkg -P aliyun-assist",
}

var Must_postClean_Data = []string{
	"/usr/local/share/aliyun*",
	"/usr/local/aegis",
	"/usr/sbin/aliyun-service",
	"/lib/systemd/system/aliyun.service",
	"/etc/systemd/system/aliyun.service",
	"/etc/init.d/aegis",
	"/usr/local/share/aliyun-assist",
	"/usr/sbin/aliyun-service",
	"/usr/sbin/aliyun_installer",
	"/usr/sbin/acs-plugin-manager",
	"/etc/systemd/system/aegis.service",
	"/lib/systemd/system/aegis.service",
	"/usr/local/cloudmonitor",
}

var Must_bannedIPs = []string{
	"140.205.201.0/24",
	"140.205.225.0/24",    // aliyun scan ip
	"140.205.205.0/24",
	"106.11.68.0/24",
	"110.75.102.0/24",
	"140.205.140.0/24",
	"10.143.23.0/24",
	"10.181.0.0/24",
	"10.181.2.0/24",
	"100.100.25.0/24",
	"110.173.196.0/24",
	"110.75.114.0/24",		// aegis server ip
	"47.110.180.128/25",		// vulnscan ip
}

//var Binaries2inst = []string{
//	"gost",
//	"besttrace",
//	"nps",   // tar.xz , extract into /usr/local/bin
//}

var Aptpkgs = []string{
	"tmux",
	"build-essential",
	"nmap",
	"socat",     // running on Ubuntu 18.04 only
	"netcat-traditional",
	"unzip",
	"p7zip",
	"zip",
	"xz-utils",
	"iptables",
}

var PreInst_Scripts = []string{
	"systemctl disable --now rsyslog",
	"systemctl disable --now atd",
	"systemctl disable --now aliyun",
	"systemctl disable --now aegis",
}