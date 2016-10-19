package main

const (
	REGISTER_TBANK1 string = `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <RegistrasiTBank xmlns="http://hackathon.bri.co.id/">
      <kodeMerchant>80090</kodeMerchant>
      <password>123456</password>
      <nohandphone>08980900120</nohandphone>
      <nama>pulsabersama</nama>
      <noktp>1671082303940006</noktp>
      <tempatLahir>Jakarta</tempatLahir>
      <tanggalLahir>19930212</tanggalLahir>
      <alamat>Slipi</alamat>
      <kota>Jakarta</kota>
      <email>zaki.afrani@gmail.com</email>
      <pekerjaan>coder</pekerjaan>
    </RegistrasiTBank>
  </soap12:Body>
</soap12:Envelope>`

	REGISTER_MERCHANT string = `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <RegistrasiMerchant xmlns="http://hackathon.bri.co.id/">
      <namaMerchant>pulsabersama</namaMerchant>
      <alamatMerchant>Slipi</alamatMerchant>
      <kota>Jakarta</kota>
      <pic>xxxx</pic>
      <noTelp>085669354499</noTelp>
      <email>zaki.afrani@gmail.com</email>
      <noRek>123456789987654</noRek>
      <password>123456</password>
    </RegistrasiMerchant>
  </soap12:Body>
</soap12:Envelope>`

	TOPUP string = `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <TopUpTBank xmlns="http://hackathon.bri.co.id/">
      <kodeMerchant>80090</kodeMerchant>
      <password>123456</password>
      <nohandphone>081210039328</nohandphone>
      <nominal>1000000</nominal>
    </TopUpTBank>
  </soap12:Body>
</soap12:Envelope>`

	TRANSFER string = `<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <TransferTBank xmlns="http://hackathon.bri.co.id/">
      <nohandphonePengirim>081210039328</nohandphonePengirim>
      <nohandphonePenerima>08980900120</nohandphonePenerima>
      <pin>133164</pin>
      <nominal></nominal>
    </TransferTBank>
  </soap12:Body>
</soap12:Envelope>`
)
