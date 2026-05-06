package kata

type Account struct {
	Email          string
	Password       string
	Profile        map[string]string
	PaymentMethods []string
	Notifications  map[string]bool
	AuditLog       []string
	Subscription   string
}

func NewAccount(email, password string) *Account {
	return &Account{email, password, map[string]string{"name": "", "bio": ""}, []string{}, map[string]bool{"email": true, "sms": false}, []string{}, "basic"}
}
func (a *Account) login(pw string) bool {
	if pw == a.Password {
		a.AuditLog = append(a.AuditLog, "login")
		return true
	}
	return false
}
func (a *Account) logout() bool {
	a.AuditLog = append(a.AuditLog, "logout")
	return true
}
func (a *Account) update_profile(n, b string) map[string]string {
	a.Profile["name"] = n
	a.Profile["bio"] = b
	return a.Profile
}
func (a *Account) change_password(o, n string) bool {
	if o == a.Password {
		a.Password = n
		return true
	}
	return false
}
func (a *Account) add_payment_method(c string) int {
	a.PaymentMethods = append(a.PaymentMethods, c)
	return len(a.PaymentMethods)
}
func (a *Account) remove_payment_method(i int) []string {
	a.PaymentMethods = append(a.PaymentMethods[:i], a.PaymentMethods[i+1:]...)
	return a.PaymentMethods
}
func (a *Account) set_notification_preference(t string, v bool) map[string]bool {
	a.Notifications[t] = v
	return a.Notifications
}
func (a *Account) export_data() map[string]any {
	return map[string]any{
		"email":          a.Email,
		"profile":        a.Profile,
		"paymentMethods": a.PaymentMethods,
		"notifications":  a.Notifications,
		"auditLog":       a.AuditLog,
		"subscription":   a.Subscription,
	}
}
func (a *Account) log_access(action string) int {
	a.AuditLog = append(a.AuditLog, action)
	return len(a.AuditLog)
}
func (a *Account) check_subscription() string {
	return a.Subscription
}
func (a *Account) upgrade_subscription(p string) string { a.Subscription = p; return a.Subscription }
