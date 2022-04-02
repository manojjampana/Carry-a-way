package Tests
import(
	"CAW/Backend/signupauth/controllers"
	"CAW/Backend/signupauth/database"
	"CAW/Backend/signupauth/models"
	
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"bytes"
	"github.com/gofiber/fiber/v2"
)
//Signup test
func TestSignup(t *testing.T){
	err:=c.Register()
	userinf:=User{
		
		"testfirstname",
		"testlastname",
		"test@a.com",
		"password",
	}
	body, err := json.Marshal(userinf)
	check(err)
	req, err := http.NewRequest("POST", "localhost:8000/api/register", bytes.NewReader(body))
    check(err)
	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.Register(c*fiber.Ctx))
	handler.ServeHTTP(rr, req)

  if status := rr.Code; status != http.StatusOK && status != http.StatusBadRequest {
    t.Errorf("handler returned wrong status code: got %v want %v or %v",
      status, http.StatusOK, http.StatusBadRequest)
  }

}
//Signin test
func TestLogin(t *testing.T){
	// setup
	err = database.DB.ConnectDB(testusername, testpassword, testaddress, testdbName)
	defer database.DB.disconnectDB()
  
	signinInfo := SigninInfo{
	  "testUsername",
	  "testPassword",
	}
  
	body, err := json.Marshal(signinInfo)
	check(err)
  
	req, err := http.NewRequest("POST", "localhost:8000/login", bytes.NewReader(body))
	check(err)
  
  
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.Login(c*fiber.Ctx))
  
	handler.ServeHTTP(rr, req)
  
	if status := rr.Code; status != http.StatusOK {
	  t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
	}
  
	
  
  }
  
func check(err error) {
	if err != nil {
	  log.Fatal(err)
	}
  }
  
