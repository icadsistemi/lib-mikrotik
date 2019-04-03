package mikrotik

import "fmt"

// ===============================
//
// Mikrotik CMDs
//
// ===============================

//Set value of item by id
func (mik *Mikrotik) Set(apipath, id string, v interface{}) error {
	args := append([]string{"=.id=" + id}, ToArgs(v)...)
	_, err := mik.RunArgs(apipath, args...)
	return err
}

//SetOne set value to one field
func (mik *Mikrotik) SetOne(apipath, name, value string) error {
	_, err := mik.RunArgs(apipath, fmt.Sprintf("=%s=%s", name, value))
	return err
}

//Remove item by id
func (mik *Mikrotik) Remove(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Enable item by id
func (mik *Mikrotik) Enable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Disable item by id
func (mik *Mikrotik) Disable(apipath, id string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id)
	return err
}

//Comment - add comment to item by id
func (mik *Mikrotik) Comment(apipath, id, comment string) error {
	_, err := mik.RunArgs(apipath, "=.id="+id, "=comment="+comment)
	return err
}

// ===============================
//
// Basic CMDs
//
// ===============================

type cmd struct {
	mikrotik *Mikrotik
	path     string
}

func (c *cmd) List(v interface{}) error {
	return c.mikrotik.Print(c.path+"/print", v)
}

func (c *cmd) Find(where string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "?"+where)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

func (c *cmd) Get(id string, v interface{}) error {
	re, err := c.mikrotik.RunArgs(c.path+"/print", "where", "?.id="+id)
	if err != nil {
		return err
	}

	return c.mikrotik.ParseResponce(re, v)
}

func (c *cmd) Add(v interface{}) error {
	return c.mikrotik.Add(c.path+"/add", v)
}

func (c *cmd) Set(id string, v interface{}) error {
	return c.mikrotik.Set(c.path+"/set", id, v)
}

func (c *cmd) Remove(id string) error {
	return c.mikrotik.Remove(c.path+"/remove", id)
}

func (c *cmd) Enable(id string) error {
	return c.mikrotik.Enable(c.path+"/enable", id)
}

func (c *cmd) Disable(id string) error {
	return c.mikrotik.Disable(c.path+"/disable", id)
}

func (c *cmd) Comment(id, comment string) error {
	return c.mikrotik.Comment(c.path+"/comment", id, comment)
}
