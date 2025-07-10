package public

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	// "github.com/jung-kurt/gofpdf"
)

/*func PrintQueuePDF(number string, serviceName string, location string) error {
	// Nama file
	timestamp := time.Now()
	fileName := fmt.Sprintf("ticket_%s.pdf", timestamp.Format("150405"))

	// Buat PDF A6
	pdf := gofpdf.New("P", "mm", "A6", "")
	pdf.AddPage()

	// Logo (kalau ada)
	if _, err := os.Stat("./assets/logo.jpg"); err == nil {
		pdf.ImageOptions("./assets/logo.jpg", 30, 5, 45, 0, false, gofpdf.ImageOptions{}, 0, "")
	}

	pdf.Ln(35)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 10, "NOMOR ANTRIAN", "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "B", 40)
	pdf.CellFormat(0, 18, number, "", 1, "C", false, 0, "")

	// Info layanan
	pdf.SetFont("Arial", "", 12)
	pdf.Ln(5)
	pdf.CellFormat(0, 8, fmt.Sprintf("Layanan : %s", serviceName), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 8, fmt.Sprintf("Waktu   : %s", timestamp.Format("2006-01-02 15:04:05")), "", 1, "", false, 0, "")
	pdf.CellFormat(0, 8, fmt.Sprintf("Lokasi  : %s", location), "", 1, "", false, 0, "")

	// QR code (optional: link ke detail antrian)
	qrData := fmt.Sprintf("https://yourdomain.com/antrian/%s", number)
	pdf.ImageOptions("https://api.qrserver.com/v1/create-qr-code/?size=80x80&data="+qrData,
		60, 105, 30, 0, false, gofpdf.ImageOptions{}, 0, "")

	pdf.Ln(30)
	pdf.SetFont("Arial", "I", 11)
	pdf.CellFormat(0, 10, "Silakan tunggu hingga dipanggil", "", 1, "C", false, 0, "")

	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		return err
	}

	// Print via Adobe Reader (silent)
	cmd := exec.Command("cmd", "/C", "start", "/min", "acrord32.exe", "/p", "/h", fileName)
	return cmd.Run()
} */

func centerText(text string, lineWidth int) string {
	textLen := len(text)
	if textLen >= lineWidth {
		return text // gak usah dipusatkan kalau lebih panjang
	}
	padding := (lineWidth - textLen) / 2
	return fmt.Sprintf("%s%s", spaces(padding), text)
}

func spaces(n int) string {
	return fmt.Sprintf("%*s", n, "")
}

func printQueueNumber(number string) error {
	location := "KELURAHAN JAKASAMPURNA"
	service := "TERIMA KASIH"
	date := time.Now().Format("02 Jan 2006 15:04")

	lineWidth := 32

	content := fmt.Sprintf(`
%s
%s

%s

%s
%s
%s
%s
%s
`,
		centerText("================================", lineWidth),
		centerText("NOMOR ANTRIAN", lineWidth),
		centerText(number, lineWidth),
		centerText("Silakan tunggu dipanggil", lineWidth),
		centerText("--------------------------------", lineWidth),
		centerText(location, lineWidth),
		centerText(service, lineWidth),
		centerText(date, lineWidth),
	)

	fileName := "ticket.txt"
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("notepad.exe", "/p", fileName)
	return cmd.Run()
}
