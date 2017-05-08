package parsers

var assetListTestCases = []Case{
	{
		"Simple",
		`Hurricane	1	Combat Battlecruiser`,
		&AssetList{
			Items: []AssetItem{AssetItem{Name: "Hurricane", Group: "Combat Battlecruiser", Quantity: 1}},
			lines: []int{0},
		},
		Input{},
		false, // This clashes with the simple contract format
	}, {
		"Typical",
		`720mm Gallium Cannon	1	Projectile Weapon	Medium	High	10 m3
Damage Control II	1	Damage Control		Low	5 m3
Experimental 10MN Microwarpdrive I	1	Propulsion Module		Medium	10 m3`,
		&AssetList{
			Items: []AssetItem{
				AssetItem{Name: "720mm Gallium Cannon", Quantity: 1, Group: "Projectile Weapon", Category: "Medium", Slot: "High", Volume: 10},
				AssetItem{Name: "Damage Control II", Quantity: 1, Group: "Damage Control", Slot: "Low", Volume: 5},
				AssetItem{Name: "Experimental 10MN Microwarpdrive I", Quantity: 1, Group: "Propulsion Module", Size: "Medium", Volume: 10},
			},
			lines: []int{0, 1, 2}},
		Input{},
		true,
	}, {
		"Full",
		`200mm AutoCannon I	1	Projectile Weapon	Module	Small	High	5 m3	1
10MN Afterburner II	1	Propulsion Module	Module	Medium	5 m3	5	2
Warrior II	9`,
		&AssetList{
			Items: []AssetItem{
				AssetItem{Name: "10MN Afterburner II", Quantity: 1, Group: "Propulsion Module", Category: "Module", Size: "Medium", MetaLevel: "5", TechLevel: "2", Volume: 5},
				AssetItem{Name: "200mm AutoCannon I", Quantity: 1, Group: "Projectile Weapon", Category: "Module", Size: "Small", Slot: "High", MetaLevel: "1", Volume: 5},
				AssetItem{Name: "Warrior II", Quantity: 9},
			},
			lines: []int{0, 1, 2},
		},
		Input{},
		true,
	}, {
		"With Volumes",
		`Sleeper Data Library	1080	Sleeper Components			10.82 m3`,
		&AssetList{
			Items: []AssetItem{AssetItem{Name: "Sleeper Data Library", Quantity: 1080, Group: "Sleeper Components", Volume: 10.82}},
			lines: []int{0},
		},
		Input{},
		true,
	}, {
		"With thousands separators",
		`Sleeper Data Library	1,080
Sleeper Data Library	1'080
Sleeper Data Library	1.080`,
		&AssetList{
			Items: []AssetItem{
				AssetItem{Name: "Sleeper Data Library", Quantity: 1080},
				AssetItem{Name: "Sleeper Data Library", Quantity: 1080},
				AssetItem{Name: "Sleeper Data Library", Quantity: 1080},
			},
			lines: []int{0, 1, 2},
		},
		Input{},
		false,
	},
}