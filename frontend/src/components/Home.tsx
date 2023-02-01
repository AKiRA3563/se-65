import { Container } from '@mui/system'
import React from 'react'

function Home() {

  return (
    <div>
      <Container maxWidth="md">
        <h1 style={{textAlign: "center"}}>ระบบยืมอุปกรณ์</h1>
        <h4><u>Requirements</u></h4>
        <p>
        ระบบการจองใช้บริการห้องของบริษัท Room Booking เป็นระบบที่ให้ผู้ใช้บริการซึ่งเป็นสมาชิกระบบจองใช้ห้องต้อง Log in เข้าระบบเพื่อทำการจองห้องต่าง ๆ 
        ที่สมาชิกต้องการ เช่น ห้องเรียน ห้องปฏิบัติการ หรือห้องอ่านหนังสือ นอกจากนี้ยังมีระบบการยืมอุปกรณ์ให้สมาชิกระบบจองใช้ห้องสามารถยืมอุปกรณ์เสริม 
        หรืออุปกรณ์เพิ่มเติมจากที่มีอยู่ในห้องโดยระบบยืมอุปกรณ์จะมีพนักงานระบบจองใช้ห้องเป็นผู้ดำเนินการบันทึกรายการยืมอุปกรณ์ให้กับสมาชิกที่มาทำการยืมอุปกรณ์ 
        ซึ่งระบบจะบันทึกข้อมูลที่เกี่ยวข้อง เช่น ชื่อสมาชิกที่ยืม ข้อมูลอุปกรณ์ จำนวน เวลา และวันที่ถูกยืมไป เป็นต้น
        </p>
        <br />
        <h4><u>User Story</u> (ระบบยืมอุปกรณ์)</h4>
        <p>
          <b>ในบทบาทของ</b>	พนักงานระบบ<br />
          <b>ฉันต้องการ</b>	ให้ระบบสามารถทำรายการการยืมอุปกรณ์<br />
          <b>เพื่อ</b>	ให้ฉันสามารถตรวจสอบได้ว่ามีอุปกรณ์ชิ้นใดถูกยืมไปแล้ว<br />
        </p>
      </Container>
    </div>
  )
}

export default Home