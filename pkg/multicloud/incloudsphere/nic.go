// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package incloudsphere

import (
	"strings"

	"yunion.io/x/onecloud/pkg/cloudprovider"
)

type SInstanceNic struct {
	cloudprovider.DummyICloudNic
	ins *SInstance

	Id              string  `json:"id"`
	AutoGenerated   bool    `json:"autoGenerated"`
	Name            string  `json:"name"`
	NolocalName     string  `json:"nolocalName"`
	InnerName       string  `json:"innerName"`
	DevName         string  `json:"devName"`
	IP              string  `json:"ip"`
	Netmask         string  `json:"netmask"`
	Gateway         string  `json:"gateway"`
	Mac             string  `json:"mac"`
	Model           string  `json:"model"`
	DeviceId        string  `json:"deviceId"`
	DeviceName      string  `json:"deviceName"`
	DeviceType      string  `json:"deviceType"`
	SwitchType      string  `json:"switchType"`
	VswitchId       string  `json:"vswitchId"`
	UplinkRate      int     `json:"uplinkRate"`
	UplinkBurst     int     `json:"uplinkBurst"`
	DownlinkRate    int     `json:"downlinkRate"`
	DownlinkBurst   int     `json:"downlinkBurst"`
	DownlinkQueue   string  `json:"downlinkQueue"`
	Enable          bool    `json:"enable"`
	Status          string  `json:"status"`
	InboundRate     float64 `json:"inboundRate"`
	OutboundRate    float64 `json:"outboundRate"`
	ConnectStatus   bool    `json:"connectStatus"`
	VMName          string  `json:"vmName"`
	VMId            string  `json:"vmId"`
	VMStatus        string  `json:"vmStatus"`
	VMTemplate      bool    `json:"vmTemplate"`
	NetworkName     string  `json:"networkName"`
	NetworkVlan     string  `json:"networkVlan"`
	VlanRange       string  `json:"vlanRange"`
	NetworkId       string  `json:"networkId"`
	NetworkType     string  `json:"networkType"`
	HostIP          string  `json:"hostIp"`
	HostStatus      string  `json:"hostStatus"`
	HostId          string  `json:"hostId"`
	DirectObjName   string  `json:"directObjName"`
	TotalOctets     float64 `json:"totalOctets"`
	TotalDropped    float64 `json:"totalDropped"`
	TotalPackets    float64 `json:"totalPackets"`
	TotalBytes      float64 `json:"totalBytes"`
	TotalErrors     float64 `json:"totalErrors"`
	WriteOctets     float64 `json:"writeOctets"`
	WriteDropped    float64 `json:"writeDropped"`
	WritePackets    float64 `json:"writePackets"`
	WriteBytes      float64 `json:"writeBytes"`
	WriteErrors     float64 `json:"writeErrors"`
	ReadOctets      float64 `json:"readOctets"`
	ReadDropped     float64 `json:"readDropped"`
	ReadPackets     float64 `json:"readPackets"`
	ReadBytes       float64 `json:"readBytes"`
	ReadErrors      float64 `json:"readErrors"`
	SecurityGroups  string  `json:"securityGroups"`
	AdvancedNetIP   string  `json:"advancedNetIp"`
	PortId          string  `json:"portId"`
	SdnVFId         string  `json:"sdnVFId"`
	OpenstackId     string  `json:"openstackId"`
	BindIPEnable    bool    `json:"bindIpEnable"`
	BindIP          string  `json:"bindIp"`
	PriorityEnabled bool    `json:"priorityEnabled"`
	NetPriority     string  `json:"netPriority"`
	VMType          string  `json:"vmType"`
	SystemVMType    string  `json:"systemVmType"`
	Dhcp            bool    `json:"dhcp"`
	DhcpIP          string  `json:"dhcpIp"`
	DhcpEnabled     bool    `json:"dhcpEnabled"`
	UsedDpdk        bool    `json:"usedDpdk"`
	Queues          int     `json:"queues"`
}

func (self *SInstanceNic) GetId() string {
	return self.Id
}

func (self *SInstanceNic) GetIP() string {
	return self.IP
}

func (self *SInstanceNic) GetMAC() string {
	return self.Mac
}

func (self *SInstanceNic) GetDriver() string {
	return strings.ToLower(self.Model)
}

func (self *SInstanceNic) InClassicNetwork() bool {
	return true
}

func (self *SInstanceNic) GetSubAddress() ([]string, error) {
	return []string{}, nil
}

func (self *SInstanceNic) GetINetworkId() string {
	return self.NetworkId
}

func (self *SInstanceNic) AssignAddress(ipAddrs []string) error {
	return cloudprovider.ErrNotSupported
}
