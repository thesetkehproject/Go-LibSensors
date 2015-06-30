package libsensors

const (
	SENSORS_API_VERSION = 0x421
	SENSORS_CHIP_NAME_PREFIX_ANY
	SENSORS_CHIP_NAME_ADDR_ANY = -1

	SENSORS_BUS_TYPE_ANY     = -1
	SENSORS_BUS_TYPE_I2C     = 0
	SENSORS_BUS_TYPE_ISA     = 1
	SENSORS_BUS_TYPE_PCI     = 2
	SENSORS_BUS_TYPE_SPI     = 3
	SENSORS_BUS_TYPE_VIRTUAL = 4
	SENSORS_BUS_TYPE_ACPI    = 5
	SENSORS_BUS_TYPE_HID     = 6
	SENSORS_BUS_NR_ANY       = -1
	SENSORS_BUS_NR_IGNORE    = -2
)

type sensors_bus_id struct {
	type int
	nr int
}

type sensors_chip_name struct {
}

func main() {}
