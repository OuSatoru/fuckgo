package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			conn.Close()
		}
	}()
	buf := make([]byte, 4096)
	_, err := conn.Read(buf)
	if err != nil {
		log.Print(err)
		return
	}
	conn.Write([]byte(`00003959<?xml version="1.0" encoding="GBK"?>
		<FORMDATA>
		  <HEAD>    <JKDM>21204</JKDM>    <FSDW>96008</FSDW>     <FSDQ>320981</FSDQ>     <JSDW>SBXT</JSDW>     <JSDQ>320981</JSDQ>     <JYDM>000000</JYDM>     <RETCODE>000000</RETCODE>     <RETINFO> </RETINFO>   </HEAD>   <DATA>     <YHLSH>96008201806050000000001</YHLSH>     <SBLSH>SBXT201806131625222145</SBLSH>     <CBDXLX>1</CBDXLX>     <CBDXBH>JSJ05066266X</CBDXBH>     <MC>\xcb\xce\xb2\xca\xc6\xbd</MC>     <ZJLX>01</ZJLX>     <ZJHM>320919196212285484</ZJHM>     <BZ> </BZ>     <ROWSET>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201407</KSNY>         <ZZNY>201506</ZZNY>         <JFDJH>201503051815466463</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>4860</JEDXZ>         <JEZDZ>4860</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>04</BXZL>         <YWLB>001</YWLB>         <KSNY>201307</KSNY>         <ZZNY>201906</ZZNY>         <JFDJH>201806110854392088</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>12164.6</JEDXZ>         <JEZDZ>12164.6</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201307</KSNY>         <ZZNY>201312</ZZNY>         <JFDJH>201806110856452091</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>2152.8</JEDXZ>         <JEZDZ>2152.8</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201901</KSNY>         <ZZNY>201906</ZZNY>         <JFDJH>201806110857032093</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>3060</JEDXZ>         <JEZDZ>3060</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201801</KSNY>         <ZZNY>201812</ZZNY>         <JFDJH>201806110857142095</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>6120</JEDXZ>         <JEZDZ>6120</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201701</KSNY>         <ZZNY>201712</ZZNY>         <JFDJH>201806110857222097</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>6120</JEDXZ>         <JEZDZ>6120</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201601</KSNY>         <ZZNY>201612</ZZNY>         <JFDJH>201806110857322099</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>5818.8</JEDXZ>         <JEZDZ>5818.8</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201507</KSNY>         <ZZNY>201512</ZZNY>         <JFDJH>201806110857432101</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>2758.8</JEDXZ>         <JEZDZ>2758.8</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>       <ROW>         <YWLSH>96008201806050000000001</YWLSH>         <BXZL>01</BXZL>         <YWLB>001</YWLB>         <KSNY>201401</KSNY>         <ZZNY>201406</ZZNY>         <JFDJH>201806110857532103</JFDJH>         <YJFS>3</YJFS>         <JFDC> </JFDC>         <JEDXZ>2152.8</JEDXZ>         <JEZDZ>2152.8</JEZDZ>         <JFYXGZ>0</JFYXGZ>         <CBDQ>320981</CBDQ>       </ROW>     </ROWSET>   </DATA> </FORMDATA>  `))
}
