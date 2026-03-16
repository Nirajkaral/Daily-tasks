package main

import (
	"fmt"
	"strings"
)

// ============================================================
// ABSTRACTION — Interface (shows only what's needed)
// ============================================================
type MedicineItem interface {
	MedicineType() string
	DisplayInfo()
	GetPrice() float64
	GetName() string
	IsInStock() bool
}

// ============================================================
// ENCAPSULATION — Base struct (private fields)
// ============================================================
type BaseMedicine struct {
	Name        string
	Company     string
	price       float64 // private
	stock       int     // private
	expiryDate  string  // private
}

// Getters — controlled access to private fields
func (b BaseMedicine) GetName() string    { return b.Name }
func (b BaseMedicine) GetPrice() float64  { return b.price }
func (b BaseMedicine) GetStock() int      { return b.stock }
func (b BaseMedicine) GetExpiry() string  { return b.expiryDate }
func (b BaseMedicine) IsInStock() bool    { return b.stock > 0 }

// Setters — controlled modification
func (b *BaseMedicine) SetPrice(price float64) {
	if price > 0 {
		b.price = price
		fmt.Printf("✅ Price updated to ₹%.2f\n", price)
	} else {
		fmt.Println("❌ Price must be positive!")
	}
}

func (b *BaseMedicine) AddStock(qty int) {
	if qty > 0 {
		b.stock += qty
		fmt.Printf("✅ Stock added! New stock: %d\n", b.stock)
	} else {
		fmt.Println("❌ Quantity must be positive!")
	}
}

func (b *BaseMedicine) ReduceStock(qty int) bool {
	if qty > b.stock {
		fmt.Printf("❌ Only %d units available!\n", b.stock)
		return false
	}
	b.stock -= qty
	return true
}

// ============================================================
// COMPOSITION — Different medicine types (Inheritance)
// ============================================================

// Tablet struct
type Tablet struct {
	BaseMedicine
	DosageMg int
}

// Syrup struct
type Syrup struct {
	BaseMedicine
	VolumeML int
}

// Injection struct
type Injection struct {
	BaseMedicine
	InjectionType string
}

// Capsule struct
type Capsule struct {
	BaseMedicine
	StrengthMg int
}

// ============================================================
// POLYMORPHISM — Same method, different behavior
// ============================================================

// Tablet methods
func (t Tablet) MedicineType() string { return "💊 Tablet" }
func (t Tablet) DisplayInfo() {
	stock := fmt.Sprintf("%d strips", t.GetStock())
	if !t.IsInStock() {
		stock = "❌ Out of Stock"
	}
	fmt.Printf(`
%s
   Name       : %s
   Company    : %s
   Dosage     : %d mg
   Price      : ₹%.2f per strip
   Stock      : %s
   Expiry     : %s
`, t.MedicineType(), t.Name, t.Company,
		t.DosageMg, t.GetPrice(), stock, t.GetExpiry())
}
func (t Tablet) GetPrice() float64 { return t.price }
func (t Tablet) GetName() string   { return t.Name }

// Syrup methods
func (s Syrup) MedicineType() string { return "🧴 Syrup" }
func (s Syrup) DisplayInfo() {
	stock := fmt.Sprintf("%d bottles", s.GetStock())
	if !s.IsInStock() {
		stock = "❌ Out of Stock"
	}
	fmt.Printf(`
%s
   Name       : %s
   Company    : %s
   Volume     : %d ml
   Price      : ₹%.2f per bottle
   Stock      : %s
   Expiry     : %s
`, s.MedicineType(), s.Name, s.Company,
		s.VolumeML, s.GetPrice(), stock, s.GetExpiry())
}
func (s Syrup) GetPrice() float64 { return s.price }
func (s Syrup) GetName() string   { return s.Name }

// Injection methods
func (i Injection) MedicineType() string { return "💉 Injection" }
func (i Injection) DisplayInfo() {
	stock := fmt.Sprintf("%d units", i.GetStock())
	if !i.IsInStock() {
		stock = "❌ Out of Stock"
	}
	fmt.Printf(`
%s
   Name       : %s
   Company    : %s
   Type       : %s
   Price      : ₹%.2f per unit
   Stock      : %s
   Expiry     : %s
`, i.MedicineType(), i.Name, i.Company,
		i.InjectionType, i.GetPrice(), stock, i.GetExpiry())
}
func (i Injection) GetPrice() float64 { return i.price }
func (i Injection) GetName() string   { return i.Name }

// Capsule methods
func (c Capsule) MedicineType() string { return "💉 Capsule" }
func (c Capsule) DisplayInfo() {
	stock := fmt.Sprintf("%d strips", c.GetStock())
	if !c.IsInStock() {
		stock = "❌ Out of Stock"
	}
	fmt.Printf(`
%s
   Name       : %s
   Company    : %s
   Strength   : %d mg
   Price      : ₹%.2f per strip
   Stock      : %s
   Expiry     : %s
`, c.MedicineType(), c.Name, c.Company,
		c.StrengthMg, c.GetPrice(), stock, c.GetExpiry())
}
func (c Capsule) GetPrice() float64 { return c.price }
func (c Capsule) GetName() string   { return c.Name }

// ============================================================
// MEDICINE STORE — manages all medicines
// ============================================================
type MedicineStore struct {
	StoreName string
	medicines []MedicineItem // private
	cart      []struct {
		item MedicineItem
		qty  int
	}
}

// Add medicine to store
func (s *MedicineStore) AddMedicine(item MedicineItem) {
	s.medicines = append(s.medicines, item)
	fmt.Printf("✅ '%s' added to store!\n", item.GetName())
}

// Show all medicines — POLYMORPHISM
func (s *MedicineStore) ShowAll() {
	fmt.Printf("\n🏥 %s — All Medicines:\n", s.StoreName)
	fmt.Println(strings.Repeat("=", 45))
	for _, m := range s.medicines {
		m.DisplayInfo() // Same method — different output per type
	}
	fmt.Println(strings.Repeat("=", 45))
}

// Search medicine
func (s *MedicineStore) Search(name string) MedicineItem {
	fmt.Printf("\n🔍 Searching: %s\n", name)
	for _, m := range s.medicines {
		if strings.EqualFold(m.GetName(), name) {
			fmt.Printf("✅ Found: %s\n", m.GetName())
			return m
		}
	}
	fmt.Printf("❌ '%s' not found!\n", name)
	return nil
}

// Add to cart
func (s *MedicineStore) AddToCart(name string, qty int) {
	var found MedicineItem
	for _, m := range s.medicines {
		if strings.EqualFold(m.GetName(), name) {
			found = m
			break
		}
	}
	if found == nil {
		fmt.Printf("❌ '%s' not found!\n", name)
		return
	}
	if !found.IsInStock() {
		fmt.Printf("❌ '%s' is out of stock!\n", name)
		return
	}
	s.cart = append(s.cart, struct {
		item MedicineItem
		qty  int
	}{found, qty})
	fmt.Printf("✅ %d x '%s' added to cart!\n", qty, name)
}

// Generate bill
func (s *MedicineStore) GenerateBill() {
	if len(s.cart) == 0 {
		fmt.Println("❌ Cart is empty!")
		return
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 45))
	fmt.Printf("   🧾 BILL — %s\n", s.StoreName)
	fmt.Println(strings.Repeat("=", 45))
	grandTotal := 0.0
	for _, c := range s.cart {
		total := c.item.GetPrice() * float64(c.qty)
		grandTotal += total
		fmt.Printf("  %-25s x%d  ₹%.2f\n",
			c.item.GetName(), c.qty, total)
	}
	fmt.Println(strings.Repeat("=", 45))
	fmt.Printf("  %-25s      ₹%.2f\n", "Grand Total", grandTotal)
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("  Get well soon! 💊")
	fmt.Println(strings.Repeat("=", 45))
}

// Show available medicines
func (s *MedicineStore) ShowAvailable() {
	fmt.Printf("\n✅ In Stock Medicines:\n")
	fmt.Println(strings.Repeat("=", 45))
	for _, m := range s.medicines {
		if m.IsInStock() {
			fmt.Printf("   %s — %s — ₹%.2f\n",
				m.MedicineType(), m.GetName(), m.GetPrice())
		}
	}
	fmt.Println(strings.Repeat("=", 45))
}

// ============================================================
// MAIN
// ============================================================
func medicine() {
	// Create store
	store := MedicineStore{StoreName: "HealthPlus Medicine Store"}

	// Create medicines
	m1 := &Tablet{
		BaseMedicine: BaseMedicine{
			Name: "Paracetamol", Company: "Sun Pharma",
			price: 25.50, stock: 100, expiryDate: "Dec 2026",
		},
		DosageMg: 500,
	}

	m2 := &Syrup{
		BaseMedicine: BaseMedicine{
			Name: "Cough Syrup", Company: "Cipla",
			price: 85.00, stock: 30, expiryDate: "Jun 2026",
		},
		VolumeML: 100,
	}

	m3 := &Injection{
		BaseMedicine: BaseMedicine{
			Name: "Insulin", Company: "Novo Nordisk",
			price: 450.00, stock: 15, expiryDate: "Mar 2026",
		},
		InjectionType: "Subcutaneous",
	}

	m4 := &Capsule{
		BaseMedicine: BaseMedicine{
			Name: "Amoxicillin", Company: "GSK",
			price: 120.00, stock: 50, expiryDate: "Sep 2026",
		},
		StrengthMg: 250,
	}

	m5 := &Tablet{
		BaseMedicine: BaseMedicine{
			Name: "Aspirin", Company: "Bayer",
			price: 15.00, stock: 0, expiryDate: "Jan 2027",
		},
		DosageMg: 75,
	}

	// Add medicines to store
	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("   🏥 HealthPlus Medicine Store")
	fmt.Println(strings.Repeat("=", 45))
	store.AddMedicine(m1)
	store.AddMedicine(m2)
	store.AddMedicine(m3)
	store.AddMedicine(m4)
	store.AddMedicine(m5)

	// Show all medicines — POLYMORPHISM
	store.ShowAll()

	// Show available medicines
	store.ShowAvailable()

	// Search medicine
	store.Search("Paracetamol")
	store.Search("Vitamin C")  // ❌ not found

	// Add to cart
	fmt.Println("\n🛒 Adding to Cart:")
	store.AddToCart("Paracetamol", 3)
	store.AddToCart("Cough Syrup", 2)
	store.AddToCart("Insulin", 1)
	store.AddToCart("Aspirin", 2)       // ❌ out of stock
	store.AddToCart("Vitamin C", 1)     // ❌ not found

	// Generate bill
	store.GenerateBill()

	// Test encapsulation — update price
	fmt.Println("\n💰 Updating Paracetamol price:")
	m1.SetPrice(30.00)
	m1.SetPrice(-10.00)  // ❌ invalid

	// Test stock management
	fmt.Println("\n📦 Adding stock for Aspirin:")
	m5.AddStock(20)
	fmt.Printf("Aspirin stock: %d\n", m5.GetStock())
}