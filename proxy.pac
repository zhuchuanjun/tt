//一个很简单的调试app pac文件
function FindProxyForURL(url, host) {
 if (dnsDomainIs(host, '39.101.133.168') || dnsDomainIs(host, 'testapi.xxcg.com') || dnsDomainIs(host, 'api.xxcg.com') || dnsDomainIs(host, 'tjapi.hlquant.com') || dnsDomainIs(host, 'api.homilychart.com')) {
            //这里填你自己的局域网ip，端口号是你在Charles里设置的 默认8888
     return "PROXY 192.168.3.213:8888; DIRECT;";
 } else {
     return "DIRECT";
 }
}