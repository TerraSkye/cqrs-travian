# CQRS Travian - Event Sourced Game Engine

A Travian 3.5 clone built with CQRS and Event Sourcing in Go.

## Events Reference

### World Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `WorldCreatedEvent` | Initializes a new game world with speed and map size settings | Admin creates a new game server |
| `TickEvent` | Advances the game simulation by one tick | Server tick timer (every few seconds) |
| `WorldEndedEvent` | Marks the world as completed with a winner | An alliance completes a World Wonder to level 100 |

### Account Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `UserAccountCreatedEvent` | Creates a new user account with credentials | User submits registration form |
| `UserAccountVerificationTokenCreatedEvent` | Generates email verification token | Account creation triggers email verification |
| `UserAccountVerifiedEvent` | Marks account as verified and active | User clicks verification link in email |

### Player Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `PlayerJoinedWorldEvent` | Player enters the game world with starting location | Verified user selects tribe and joins world |
| `PlayerTribeSelectedEvent` | Records the tribe choice (Roman/Gaul/Teuton) | Player selects tribe during registration |
| `PlayerDeletedEvent` | Removes player from the game world | Player requests account deletion or inactivity timeout |
| `GoldPurchasedEvent` | Adds premium gold currency to account | Player completes gold purchase transaction |
| `GoldSpentEvent` | Deducts gold for premium features | Player uses gold for instant complete, NPC trade, etc. |
| `GoldCompletionEvent` | Instantly completes a building in queue | Player spends gold to skip building timer |
| `PlusAccountActivatedEvent` | Enables Travian Plus premium features | Player purchases Plus account with gold |
| `PlayerSittingEnabledEvent` | Allows another player to manage account | Player adds a sitter in settings |
| `PlayerSittingDisabledEvent` | Removes sitter access from account | Player removes sitter or sitter period expires |

### Village Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `VillageCreatedEvent` | Creates initial village for new player | Player joins world and spawns at starting location |
| `VillageFoundedEvent` | Establishes new village using settlers | Settlers arrive at empty valley and settle |
| `VillageConqueredEvent` | Transfers village ownership to attacker | Chief/Senator/Chieftain reduces loyalty to 0 |
| `VillageRenamedEvent` | Changes village display name | Player edits village name in village overview |
| `VillageAbandonedEvent` | Removes village from player control | Player deliberately abandons non-capital village |
| `CapitalChangedEvent` | Designates a different village as capital | Player changes capital in Palace/Residence |

### Building Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `BuildingQueuedEvent` | Adds building upgrade to construction queue | Player clicks upgrade on a building slot |
| `BuildingStartedEvent` | Begins construction with timer countdown | Queue position becomes first or slot becomes free |
| `BuildingCompletedEvent` | Finishes construction, increases building level | Construction timer reaches zero |
| `BuildingDestroyedEvent` | Reduces building level from enemy attack | Catapults hit the targeted building |
| `BuildingDemolishedEvent` | Player intentionally lowers building level | Player uses Main Building demolish feature |
| `BuildingCancelledEvent` | Removes queued building, refunds resources | Player cancels construction in queue |

### Resource Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `ResourcesProducedEvent` | Generates resources based on production rate | Server tick calculates hourly production |
| `ResourcesConsumedEvent` | Deducts resources for construction/training | Building upgrade, troop training, or celebrations |
| `ResourcesTransferredEvent` | Sends resources via merchants to another village | Player sends resources through marketplace |
| `ResourcesReceivedEvent` | Receives resources from trade or raid return | Merchants arrive with traded/raided goods |
| `ResourcesRaidedEvent` | Resources stolen by attacking troops | Successful raid calculates loot based on carry capacity |
| `ResourcesOverflowEvent` | Resources lost due to storage capacity limit | Production exceeds warehouse/granary capacity |

### Troop Training Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `TroopTrainingQueuedEvent` | Adds troops to training queue | Player orders troops in Barracks/Stable/Workshop |
| `TroopTrainingStartedEvent` | Begins training with countdown timer | Training queue position becomes active |
| `TroopTrainingCompletedEvent` | Troops added to village garrison | Training timer completes for troop batch |

### Troop Movement Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `TroopsSentEvent` | Dispatches troops to target coordinates | Player sends attack/raid/reinforcement |
| `TroopsArrivedEvent` | Troops reach destination village/oasis | Movement timer completes |
| `TroopsReturnedEvent` | Troops return home after mission | Return timer completes after attack/raid |
| `TroopsKilledEvent` | Troops lost in battle | Combat resolution calculates casualties |
| `TroopsDisbandedEvent` | Player intentionally removes troops | Player dismisses troops to reduce crop consumption |
| `ReinforcementArrivedEvent` | Allied troops arrive to defend village | Reinforcement movement completes |
| `ReinforcementRecalledEvent` | Owner recalls reinforcements home | Player recalls troops from allied village |

### Research Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `ResearchStartedEvent` | Begins researching a new troop type | Player researches unit in Academy |
| `ResearchCompletedEvent` | Unlocks troop type for training | Research timer completes |
| `TroopUpgradeStartedEvent` | Begins upgrading troop attack/defense | Player upgrades unit in Blacksmith/Armoury |
| `TroopUpgradeCompletedEvent` | Improves troop stats by upgrade level | Upgrade timer completes |

### Combat Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `AttackLaunchedEvent` | Records outgoing attack with troop composition | Player sends attack/raid with target selection |
| `BattleResolvedEvent` | Calculates combat outcome, casualties, loot | Attacking troops arrive at target |
| `SpyMissionCompletedEvent` | Reports scouting results (resources/troops/defenses) | Scout units complete spy mission |
| `TrapTriggeredEvent` | Gaul trapper captures attacking troops | Attacking troops fall into traps |
| `TrapReleasedEvent` | Trapped troops freed or killed | Trap owner releases or successful attack frees them |
| `LoyaltyChangedEvent` | Village loyalty reduced by chief unit | Senator/Chief/Chieftain attacks village |
| `WallDamagedEvent` | Wall level reduced by battering rams | Rams survive battle and hit the wall |

### Hero Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `HeroCreatedEvent` | Spawns hero unit for player | Player creates first village (hero auto-created) |
| `HeroLeveledUpEvent` | Hero gains level and skill points | Hero gains enough experience points |
| `HeroAttributeAssignedEvent` | Distributes skill points to hero stats | Player assigns points to strength/off/def/resource |
| `HeroAdventureStartedEvent` | Hero departs for adventure location | Player sends hero on adventure |
| `HeroAdventureCompletedEvent` | Hero returns with rewards or injuries | Hero arrives at adventure point |
| `HeroEquipmentChangedEvent` | Hero equips/unequips an item | Player changes hero inventory |
| `HeroDiedEvent` | Hero killed in battle or adventure | Hero health reaches zero |
| `HeroRevivedEvent` | Hero regenerates after death | Hero revive timer completes in village |
| `HeroHealthChangedEvent` | Hero health increases or decreases | Combat damage, regeneration, or healing items |
| `AdventureSpawnedEvent` | New adventure appears for player | Server spawns adventures based on hero level |

### Alliance Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `AllianceCreatedEvent` | Forms new alliance with founder as leader | Player creates alliance in Embassy |
| `AllianceDisbandedEvent` | Dissolves alliance, members become unaffiliated | Leader disbands or all members leave |
| `AllianceInvitationSentEvent` | Sends membership invite to a player | Alliance recruiter invites player |
| `AllianceInvitationAcceptedEvent` | Player accepts and joins alliance | Invited player clicks accept |
| `AllianceInvitationDeclinedEvent` | Player rejects alliance invitation | Invited player clicks decline |
| `AllianceMemberJoinedEvent` | Adds player to alliance roster | Invitation accepted or direct join |
| `AllianceMemberLeftEvent` | Player voluntarily leaves alliance | Player clicks leave alliance |
| `AllianceMemberKickedEvent` | Leader/deputy removes member forcibly | Alliance leader kicks member |
| `AllianceRoleChangedEvent` | Member promoted/demoted in alliance | Leader changes member role |
| `AllianceDiplomacyChangedEvent` | Sets NAP/Ally/War status with another alliance | Diplomat updates diplomacy settings |
| `AllianceDescriptionUpdatedEvent` | Changes alliance profile/description | Leader edits alliance description |

### Marketplace Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `TradeOfferCreatedEvent` | Posts trade offer on marketplace | Player creates offer to exchange resources |
| `TradeOfferAcceptedEvent` | Another player accepts trade offer | Player accepts visible marketplace offer |
| `TradeOfferCancelledEvent` | Trade offer withdrawn by creator | Player cancels their own offer |
| `TradeOfferExpiredEvent` | Trade offer times out without acceptance | Offer duration expires |
| `MerchantsSentEvent` | Merchants depart carrying resources | Player sends resources or trade accepted |
| `MerchantsArrivedEvent` | Merchants deliver resources to destination | Merchant travel timer completes |
| `MerchantsReturnedEvent` | Merchants return home available for new trades | Return journey completes |
| `NPCTradeCompletedEvent` | Instant resource ratio exchange for gold | Player uses NPC merchant with gold |

### Message Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `MessageSentEvent` | Player sends private message to another | Player submits message form |
| `MessageReadEvent` | Recipient opens/reads message | Player views message in inbox |
| `MessageDeletedEvent` | Message removed from mailbox | Player deletes message |
| `ReportCreatedEvent` | Battle/trade/adventure report generated | Action completes that generates a report |
| `ReportReadEvent` | Player views generated report | Player opens report in reports tab |
| `ReportDeletedEvent` | Report removed from player's reports | Player deletes report |
| `IGMSentEvent` | Alliance-wide message broadcast | Alliance leader/diplomat sends IGM |

### Oasis Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `OasisConqueredEvent` | Player claims oasis for resource bonus | Hero clears oasis and village has Hero's Mansion slot |
| `OasisLostEvent` | Oasis ownership transfers to another player | Another player conquers the oasis |
| `OasisAbandonedEvent` | Player releases oasis voluntarily | Player abandons oasis in Hero's Mansion |
| `OasisAnimalsSpawnedEvent` | Wild animals regenerate in unoccupied oasis | Server tick spawns animals in oasis |
| `OasisRaidedEvent` | Resources and animals taken from oasis | Player raids oasis for resources |

### World Wonder & Endgame Events

| Event | Description | Trigger |
|-------|-------------|---------|
| `WorldWonderVillageClaimedEvent` | Alliance claims WW village site | Player conquers Natar WW village |
| `WorldWonderLevelIncreasedEvent` | World Wonder construction advances | WW building upgrade completes |
| `WorldWonderDestroyedEvent` | World Wonder levels destroyed by attack | Catapults hit World Wonder building |
| `WorldWonderCompletedEvent` | World Wonder reaches level 100, game ends | Final WW construction completes |
| `ArtifactSpawnedEvent` | Artifacts appear in Natar villages | Server triggers artifact spawn day |
| `ArtifactCapturedEvent` | Artifact stolen by attacking player | Hero captures artifact from village |
| `ArtifactActivatedEvent` | Artifact bonus becomes active | 24-hour activation timer completes |
| `NatarVillageSpawnedEvent` | NPC Natar village appears on map | Server spawns Natar villages for endgame |

## Supporting Types

Located in `engine/support/`:

| File | Types | Description |
|------|-------|-------------|
| `building.go` | `BuildingType` | All 44 building types (resource fields + village buildings) |
| `resource.go` | `Resources` | [4]int64 array for wood, clay, iron, crop |
| `tribe.go` | `Tribe` | Roman, Gaul, Teuton |
| `troop.go` | `TroopType`, `Troops`, `MovementType` | All unit types, troop maps, movement purposes |
| `hero.go` | `HeroAttribute`, `EquipmentSlot`, `ItemType` | Hero stats, equipment slots, item definitions |
| `alliance.go` | `AllianceRole`, `DiplomacyType` | Member roles, diplomatic relations |
| `report.go` | `ReportType` | Battle, reinforce, scout, adventure, trade, other |
| `oasis.go` | `OasisType` | Oasis bonus types (25%/50% resource bonuses) |
