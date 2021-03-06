package constants

const (
	MOB_STATUS_WATK = 0x01
	MOB_STATUS_WDEF = 0x02
	MOB_STATUS_MATK = 0x04
	MOB_STATUS_MDEF = 0x08
	MOB_STATUS_ACC  = 0x10

	MOB_STATUS_AVOID  = 0x20
	MOB_STATUS_SPEED  = 0x40
	MOB_STATUS_STUN   = 0x80
	MOB_STATUS_FREEZE = 0x100
	MOB_STATUS_POISON = 0x200

	MOB_STATUS_SEAL              = 0x400
	MOB_STATUS_NO_CLUE1          = 0x800
	MOB_STATUS_WEAPON_ATTACK_UP  = 0x1000
	MOB_STATUS_WEAPON_DEFENSE_UP = 0x2000
	MOB_STATUS_MAGIC_ATTACK_UP   = 0x4000

	MOB_STATUS_MAGIC_DEFENSE_UP = 0x8000
	MOB_STATUS_DOOM             = 0x10000
	MOB_STATUS_SHADOW_WEB       = 0x20000
	MOB_STATUS_WEAPON_IMMUNITY  = 0x40000
	MOB_STATUS_MAGIC_IMMUNITY   = 0x80000

	MOB_STATUS_NO_CLUE2        = 0x100000
	MOB_STATUS_NO_CLUE3        = 0x200000
	MOB_STATUS_NINJA_AMBUSH    = 0x400000
	MOB_STATUS_NO_CLUE4        = 0x800000
	MOB_STATUS_VENOMOUS_WEAPON = 0x1000000

	MOB_STATUS_NO_CLUE5              = 0x2000000
	MOB_STATUS_NO_CLUE6              = 0x4000000
	MOB_STATUS_EMPTY                 = 0x8000000 // All mobs have this when they spawn
	MOB_STATUS_HYPNOTIZE             = 0x10000000
	MOB_STATUS_WEAPON_DAMAGE_REFLECT = 0x20000000

	MOB_STATUS_MAGIC_DAMAGE_REFLECT = 0x40000000
	MOB_STATUS_NO_CLUE7             = 0x80000000 // Last bit you can use with 4 bytes

	PLAYER_STATUS_CURSE       = 0x01
	PLAYER_STATUS_WEAKNESS    = 0x02
	PLAYER_STATUS_DARKNESS    = 0x04
	PLAYER_STATUS_SEAL        = 0x08
	PLAYER_STATUS_POISON      = 0x10
	PLAYER_STATUS_STUN        = 0x20
	PLAYER_STATUS_SLOW        = 0x40
	PLAYER_STATUS_SEDUCE      = 0x80
	PLAYER_STATUS_ZOMBIFY     = 0x100
	PLAYER_STATUS_CRAZY_SKULL = 0x200
)
