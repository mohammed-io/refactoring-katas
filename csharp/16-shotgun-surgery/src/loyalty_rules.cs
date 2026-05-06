public static class LoyaltyRules
{
    public static decimal get_discount_for_tier(string t)
    {
        if (t == "bronze")
            return 0.05m;
        if (t == "silver")
            return 0.1m;
        if (t == "gold")
            return 0.15m;
        if (t == "platinum")
            return 0.2m;
        return 0;
    }

    public static string get_label_for_tier(string t)
    {
        if (t == "bronze")
            return "Bronze Member";
        if (t == "silver")
            return "Silver Member";
        if (t == "gold")
            return "Gold Member";
        if (t == "platinum")
            return "Platinum Member";
        return "Standard";
    }

    public static int get_threshold_for_tier(string t)
    {
        if (t == "bronze")
            return 100;
        if (t == "silver")
            return 500;
        if (t == "gold")
            return 2000;
        if (t == "platinum")
            return 10000;
        return 0;
    }

    public static string get_color_for_tier(string t)
    {
        if (t == "bronze")
            return "#CD7F32";
        if (t == "silver")
            return "#C0C0C0";
        if (t == "gold")
            return "#FFD700";
        if (t == "platinum")
            return "#E5E4E2";
        return "#000000";
    }

    public static decimal calculate_total(decimal s, string t)
    {
        return s * (1 - get_discount_for_tier(t));
    }
}
