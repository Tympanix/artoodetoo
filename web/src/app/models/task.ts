import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Subject } from 'rxjs/Subject';
import { Unit, IUnit, Model, Event } from '../model'
import * as _ from "lodash";

export interface ITask {
  name: string
  event: IUnit
  actions: IUnit[]
  // running: boolean
}

export class Task implements ITask, Model {
  // Model properties
  name: string = ""
  event: Unit = null
  actions: Unit[] = []
  running: boolean = false

  private eventRef: string

  // State properties
  units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)

  constructor(fields?: {
    name?: string
  }) {
    if (fields) Object.assign(this, fields)
  }

  static fromJson(model: ITask): Task {
    let task = new Task()
    Object.assign(task, model)
    task.eventRef = model.event as Object as string
    //task.event = Unit.fromJson(model.event)
    task.actions = model.actions.map(action => Unit.fromJson(action))
    task.updateUnitList()
    return task
  }

  resolveEvent(events: Event[]) {
    let event = events.find(e => e.uuid == this.eventRef)

    if (!event) {
      throw new Error(`Event ${this.eventRef} not found for ${this.name}`)
    }

    this.event = event
  }

  copy(): Task {
    let copy = _.cloneDeep(this)
    return copy
  }

  public toJson(): ITask {
    return {
      name: this.name,
      event: this.event.toJson(),
      actions: this.actions.map(a => a.toJson())
    }
  }

  addAction(unit: Unit) {
    this.actions.push(unit)
    this.updateUnitList()
  }

  updateUnitList() {
    let units: Unit[] = []
    this.event && units.push(this.event)
    this.actions.forEach(action => {
      action && units.push(action)
    })
    this.units.next(units)
  }

  deleteUnit(unit: Unit) {
    console.log("Deleting unit: ", unit)
    this.actions = this.actions.filter(u => u !== unit)
    this.updateUnitList()
  }

  deleteEvent(){
    console.log("Deleting event")
    this.event = null
  }

  private swapActions(indexFrom: number, indexTo: number) {
    if (indexFrom < 0 || indexFrom >= this.actions.length) return
    if (indexTo < 0 || indexTo >= this.actions.length) return
    let temp = this.actions[indexTo]
    this.actions[indexTo] = this.actions[indexFrom]
    this.actions[indexFrom] = temp
  }

  moveUnitUp(unit: Unit) {
    let indexFrom = this.actions.indexOf(unit)
    this.swapActions(indexFrom, indexFrom - 1)
  }

  moveUnitDown(unit: Unit) {
    let indexFrom = this.actions.indexOf(unit)
    this.swapActions(indexFrom, indexFrom + 1)
  }

  getSourceByName(name: string): Unit {
    if (this.event && this.event.name == name) return this.event
    return this.actions.find(unit => unit.name == name)
  }

}
