package database

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

type Supabase struct {
	SupabaseUrl string
	SupabaseKey string
}

func NewSupabase() *supa.Client {

	supabaseAuth := Supabase{
		SupabaseUrl: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_PRIVATE_KEY"),
	}
	supabaseClient := supa.CreateClient(supabaseAuth.SupabaseUrl, supabaseAuth.SupabaseKey)

	return supabaseClient
}
