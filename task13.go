package main

import "fmt"

type administrator interface {
	administrate(system string)
}

type developer interface {
	develop(system string)
}

type adminlist struct {
	list []administrator
}

func (l *adminlist) Enqueue(a administrator) {
	l.list = append(l.list, a)
}

func (l *adminlist) Dequeue() administrator {
		a := l.list[0]
		l.list = l.list[1:]
		return a
}

type devlist struct {
	list []developer
}

func (l *devlist) Enqueue(d developer) {
	l.list = append(l.list, d)
}

func (l *devlist) Dequeue() developer {
	d := l.list[0]
	l.list = l.list[1:]
	return d
}

type sysadmin struct {
	name string
}

func (s sysadmin) administrate(system string) {
	fmt.Printf("Sysadmin %s is administering %s\n", s.name, system)
}

type programmer struct {
	name string
}

func (p programmer) develop(system string){
	fmt.Printf("programmer %s is developing %s\n", p.name, system)
}

type company struct {
	administrator
	developer
}

func main(){
	var admins adminlist

	var devs devlist

	admins.Enqueue(sysadmin{name: "Alice"})

	devs.Enqueue(programmer{name: "Bob"})
	devs.Enqueue(programmer{name: "Den"})

	cmp := company{
		administrator: admins.Dequeue(),
		developer: devs.Dequeue(),
	}

	admins.Enqueue(cmp)
	devs.Enqueue(cmp)

	tasks := []struct {
		needsAdmin bool
		system string
	}{
		{needsAdmin: false, system: "xenia"},
		{needsAdmin: true, system: "pillar"},
		{needsAdmin: false, system: "omega"},
	}

	for _, task := range tasks {
		if task.needsAdmin {
			admin := admins.Dequeue()
			admin.administrate(task.system)
			continue
		}

		dev := devs.Dequeue()
		dev.develop(task.system)
	}
}
