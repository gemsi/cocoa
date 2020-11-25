package appkit

// These constants are used as a bitmask for specifying a set of menu or menu item properties, and are contained by the propertiesToUpdate property.
type MenuProperties uint

const (
	MenuPropertyItemTitle                    MenuProperties = 1 << 0 // the menu item's title
	MenuPropertyItemAttributedTitle          MenuProperties = 1 << 1 // the menu item's attributed title
	MenuPropertyItemKeyEquivalent            MenuProperties = 1 << 2 // the menu item's key equivalent
	MenuPropertyItemImage                    MenuProperties = 1 << 3 // the menu item's image
	MenuPropertyItemEnabled                  MenuProperties = 1 << 4 // whether the menu item is enabled or disabled
	MenuPropertyItemAccessibilityDescription MenuProperties = 1 << 5 // the menu item's accessibility description
)
