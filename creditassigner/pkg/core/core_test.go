package core

func testCore(t *testing.T) {
	c := CreditAssigner{}

	a,b,c,e := c.Assign(3100)

	if err == nil {
		t.Fatal("Tendria que dar error ya que no es posible")
	}

	a,b,c,e := c.Assign(300)

	if err != nil {
		t.Fatal("No tendria que dar error ya que es posible")
	}
}