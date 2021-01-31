package main

import (
	"fmt"

	"github.com/kcasctiv/amichan"
)

func main() {
	port := 7080
	keepalive := true
 	user := "asterisk"
  	pass :=  "asterisk"
  	ip := "127.0.0.1"
	//יצירת חיבור לAMI
	a := amichan.New(user,pass,ip , port, keepalive)
	a.Connect()

	for {
		//לולאה אינסופית שמאזינה לAMI
		select {
		case err := <-a.Err():
		// להדפיס לקונסול במקרה של שגיאה	
			fmt.Println(err)
		case event := <-a.Event():
			// במקרה של אירוע לבדוק האם קיבלתי אירוע של הקשות
			if event.Name() == "DTMFEnd" {
				
				// בדיקה האם קיים רשומה באובייקט של האירוע
				// הפונקציה מחזירה שתי ערכים
				//1. האם האובייקט קיים
				// 2. את הערך באובייקט
				// כיוון שאסור (לפעמים) להשתמש במשתנה ריק
				// אז אני קודם בודק האם האובייקט קיים
				// ורק אז משתמש בו
				
				// אז כאן אני מדפיס את המספר שלך המאזין
				// ואת המקש שהוא הקיש
				// לא זוכר מה עשיתי עם הערך השלישי
				// כמובן שבמקום להדפיס אתה יכול לשלוח אותו למקום אחר
				// למשל API
				Num, ok := event.Field("CallerIDNum")
				if ok {
					fmt.Println(Num)
				}
        
				Digit, ok := event.Field("Digit")
				if ok {
					fmt.Println(Digit)
				}
        
				ID, ok := event.Field("Uniqueid")
				if ok {
					fmt.Println(ID)
				}
				
			}
		}
	}
}
