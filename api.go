package gos7

// Copyright 2018 Trung Hieu Le. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.
import (
	"time"
)

// Client interface s7 client
type Client interface {

	//AGReadDB Read data blocks from PLC
	AGReadDB(dbNumber int, start int, size int, buffer []byte) (err error)
	// AGWriteDB write data blocks into PLC
	AGWriteDB(dbNumber int, start int, size int, buffer []byte) (err error)
	// AGReadMB Read Merkers area from PLC
	AGReadMB(start int, size int, buffer []byte) (err error)
	// AGWriteMB Write Merkers from into PLC
	AGWriteMB(start int, size int, buffer []byte) (err error)
	// AGReadEB Read IPI from PLC
	AGReadEB(start int, size int, buffer []byte) (err error)
	// AGWriteEB Write IPI into PLC
	AGWriteEB(start int, size int, buffer []byte) (err error)
	// AGReadAB Read IPU from PLC
	AGReadAB(start int, size int, buffer []byte) (err error)
	// AGWriteAB Write IPU into PLC
	AGWriteAB(start int, size int, buffer []byte) (err error)
	// AGReadTM Read timer from PLC
	AGReadTM(start int, size int, buffer []byte) (err error)
	// AGWriteTM Write timer into PLC
	AGWriteTM(start int, size int, buffer []byte) (err error)
	// AGReadCT Read counter from PLC
	AGReadCT(start int, size int, buffer []byte) (err error)
	// AGWriteCT Write counter into PLC
	AGWriteCT(start int, size int, buffer []byte) (err error)
	// AGReadMulti multi read area
	AGReadMulti(dataItems []S7DataItem, itemsCount int) (err error)
	// AGWriteMulti multi write area
	AGWriteMulti(dataItems []S7DataItem, itemsCount int) (err error)
	// DBFill /*block*/
	DBFill(dbnumber int, fillchar int) error
	DBGet(dbnumber int, usrdata []byte, size int) error
	// Read general read function with S7 sytax
	Read(variable string, buffer []byte) (value interface{}, err error)
	// GetAgBlockInfo Get block  infor in AG area, refer an S7BlockInfor pointer
	GetAgBlockInfo(blocktype int, blocknum int) (info S7BlockInfo, err error)

	// PLCHotStart Hotstart PLC, Puts the CPU in RUN mode performing an HOT START.
	PLCHotStart() error
	// PLCColdStart Cold start PLC, change CPU into runmode performing and COLD START
	PLCColdStart() error
	// PLCStop change CPU to stop mode
	PLCStop() error
	// PLCGetStatus return CPU status: running/stopped
	PLCGetStatus() (status int, err error)
	// PGListBlocks list all blocks in PLC, return a Blockslist which contains list of OB, DB, ...
	PGListBlocks() (list S7BlocksList, err error)
	// SetSessionPassword set the session password for PLC to meet its security level
	SetSessionPassword(password string) error
	// ClearSessionPassword clear the password set for current session
	ClearSessionPassword() error
	// GetProtection return the CPU protection level info, refer to: §33.19 of "System Software for S7-300/400 System and Standard Functions"
	//return S7Protection and its properties.
	GetProtection() (protection S7Protection, err error)
	// GetOrderCode get CPU order code, return S7OrderCode
	GetOrderCode() (info S7OrderCode, err error)
	// GetCPUInfo get CPU info, return S7CpuInfo and its properties
	GetCPUInfo() (info S7CpuInfo, err error)
	// GetCPInfo get CP info, return S7CpInfo and its properties
	GetCPInfo() (info S7CpInfo, err error)
	// PGClockRead /*datetime*/
	//read clock on PLC, return a time
	PGClockRead(datetime time.Time) error
	// PGClockWrite write clock to PLC with datetime input
	PGClockWrite() (dt time.Time, err error)

	IsConnected() bool
}
