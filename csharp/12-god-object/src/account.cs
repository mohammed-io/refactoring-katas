public class Account
{
    public string Email;
    public string Password;

    public Dictionary<string, string> Profile = new()
    {
        { "name", "" },
        { "bio", "" }
    };

    public List<string> PaymentMethods = new();

    public Dictionary<string, bool> Notifications = new()
    {
        { "email", true },
        { "sms", false }
    };

    public List<string> AuditLog = new();
    public string Subscription = "basic";

    public Account(string e, string p)
    {
        Email = e;
        Password = p;
    }

    public bool login(string pw)
    {
        if (pw == Password)
        {
            AuditLog.Add("login");
            return true;
        }
        return false;
    }

    public Dictionary<string, string> update_profile(string n, string b)
    {
        Profile["name"] = n;
        Profile["bio"] = b;
        return Profile;
    }

    public bool change_password(string o, string n)
    {
        if (o == Password)
        {
            Password = n;
            return true;
        }
        return false;
    }

    public int add_payment_method(string c)
    {
        PaymentMethods.Add(c);
        return PaymentMethods.Count;
    }

    public List<string> remove_payment_method(int i)
    {
        PaymentMethods.RemoveAt(i);
        return PaymentMethods;
    }

    public Dictionary<string, bool> set_notification_preference(string t, bool v)
    {
        Notifications[t] = v;
        return Notifications;
    }

    public string upgrade_subscription(string p)
    {
        Subscription = p;
        return Subscription;
    }
}
