# Canonical Naming Mapping for Refactoring Katas

## Naming Standard
- **Class names**: PascalCase (same across all languages)
- **Method names**: snake_case (same across all languages)
- **File base names**: snake_case (same across all languages)
- Go uses structs instead of classes, but same names

## Kata-by-Kata Mapping

| Kata | Canonical Class | Canonical File | Primary Method | Secondary Methods |
|------|----------------|----------------|----------------|-------------------|
| 01-magic-receipt | Receipt | receipt | calculate_total | |
| 02-never-ending-function | OrderProcessor | order_processor | process_order | |
| 03-copy-paste-payroll | PayrollCalculator | payroll_calculator | generate_payslips | |
| 04-cryptic-calculator | Calculator | calculator | normalize | |
| 05-zombie-code | CampaignSender | campaign_sender | send_campaign | |
| 06-nested-if-hell | PackageValidator | package_validator | can_deliver | |
| 07-switch-factory | PlaylistCurator | playlist_curator | create_playlist | |
| 08-jealous-function | KitchenTicket | kitchen_ticket | print_ticket | |
| 09-stringly-typed | IssueUpdater | issue_updater | update_issue | |
| 10-null-spiral | RoomLookup | room_lookup | get_room_for_student | |
| 11-parameter-soup | TripBooking | trip_booking | book_trip | |
| 12-god-object | Account | account | login, logout, update_profile, change_password, add_payment_method, remove_payment_method, set_notification_preference, export_data, log_access, check_subscription, upgrade_subscription | |
| 13-chatterbox | InventoryItem | inventory_item | get_id, get_name, get_batch_number, get_cache_timestamp, get_row_id, get_quantity | |
| 14-messenger | NotificationClient | notification_client | send | |
| 15-inheritance-abuse | Employee | employee | calculate_bonus | |
| 16-shotgun-surgery | LoyaltyRules | loyalty_rules | get_discount_for_tier, get_label_for_tier, get_threshold_for_tier, get_color_for_tier | |
| 17-taxonomy-trap | Vehicle | vehicle | daily_rate, fuel_cost | |
| 18-big-bang-algorithm | FraudDetector | fraud_detector | detect | |
| 19-tangled-web | ConfigLoader | config_loader | load_config | |
| 20-legacy-monolith | LegacySystem | legacy_system | process_everything | |

## Language-Specific Notes

### JavaScript
- Convert all functions to classes
- Use snake_case method names (unconventional but consistent)
- Export class: `export { Receipt }`
- Constructor for initialization

### Python
- Convert all functions to classes
- Use snake_case method names (conventional)
- `__init__` for initialization
- `self` as first parameter

### Ruby
- Already uses classes, just verify names match
- Use snake_case method names (conventional)
- `initialize` for initialization

### Go
- Use structs with methods (no classes)
- Use snake_case method names (unconventional but consistent)
- Constructor function: `NewReceipt()` etc.
- Tests in same package (package `kata`) to access unexported methods

### C#
- Already uses classes, rename methods to snake_case (unconventional but consistent)
- Keep using `public class` and `public` methods
- Use `record` types where appropriate for data structures
