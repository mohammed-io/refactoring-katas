# Refactor Python Katas to Classes

## Status: in_progress

## Context
Refactor all 20 Python refactoring katas to use classes instead of standalone functions, following canonical names from CANONICAL_NAMES.md.

## Value Proposition
- All Python source files use classes with canonical names
- Test files updated to instantiate classes and call methods
- All tests pass after refactoring
- Code logic preserved (still smelly)

## Todos
- [x] Kata 01: Receipt class (already correct, update tests)
- [x] Kata 02: OrderProcessor class, rename file
- [x] Kata 03: PayrollCalculator class, rename file
- [x] Kata 04: Calculator class, rename file
- [x] Kata 05: CampaignSender class, rename file
- [x] Kata 06: LoanApprover class, rename file
- [x] Kata 07: PlaylistCurator class, rename file
- [x] Kata 08: KitchenTicket class, convert Order+function to class method
- [x] Kata 09: IssueUpdater class, rename file
- [x] Kata 10: RoomLookup class, rename file
- [x] Kata 11: TripBooking class (already correct file, update tests)
- [x] Kata 12: Account class (verify already correct)
- [x] Kata 13: InventoryItem class (verify already correct)
- [x] Kata 14: NotificationClient class, rename file
- [x] Kata 15: Employee class, rename file
- [x] Kata 16: LoyaltyRules class, rename file
- [x] Kata 17: Vehicle class, rename file
- [x] Kata 18: FraudDetector class, wrap function in class
- [x] Kata 19: ConfigLoader class, wrap function in class
- [x] Kata 20: LegacySystem class, rename file
- [x] Run all tests and verify they pass

## Notes
- Keep code smelly - only structural changes
- Methods must be snake_case
- Files must match canonical snake_case names
- Classes must be PascalCase
