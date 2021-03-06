import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router'
import { Task, Unit } from '../../model'
import { ApiService } from '../../api.service'
import { MdDialog, MdDialogRef, MdSnackBar } from '@angular/material';
import { UnitDialog, TaskDialog, EventDialog, CycleDialog } from '../../dialogs'

@Component({
  selector: 'taskeditor',
  templateUrl: './taskeditor.component.html',
  styles: []
})
export class TaskeditorComponent implements OnInit {
  @Input() task: Task

  events: Unit[]

  constructor(private api: ApiService, public dialog: MdDialog, private snackBar: MdSnackBar, private router: Router) { }

  ngOnInit() {
    this.api.events.subscribe(e => this.events = e)
  }

  eventChange(){
    this.task.updateUnitList()
  }

  createTask(): void {
    this.api.createTask(this.task).subscribe(() => {
      this.snackBar.open(this.task.name + " has been created", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      this.gotoDashboard()
    })
  }

  updateTask() {
    this.api.updateTask(this.task).subscribe(() => {
      this.snackBar.open(this.task.name + " has been saved", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      this.gotoDashboard()
    })

  }

  gotoDashboard() {
    this.router.navigateByUrl("/dashboard")
  }

  deleteTask(){
    this.api.deleteTask(this.task).subscribe()
  }

  submitTask() {
    CycleDialog.check(this.dialog, this.snackBar, this.task).then(() => {
      if (this.task.isSaved) {
        this.updateTask()
      } else {
        this.createTask()
      }
    })
  }

  openUnitDialog() {
    let dialogRef = this.dialog.open(UnitDialog, {
      height: '500px',
      width: '750px',
    });
    dialogRef.afterClosed().subscribe(unit => {
      if (unit) {
          this.task.addAction(unit);
      }
    });
  }

  openTaskDialog(){
    let dialogRef = this.dialog.open(TaskDialog, {
      width: '600px'
    });

    dialogRef.afterClosed().subscribe(name => {
      if(name != undefined && name != ""){
        this.task = new Task({name: name})
      }
    })
  }

  openEventDialog(){
    let dialogRef = this.dialog.open(EventDialog, {
      height: '500px',
      width: '750px',
    })

    dialogRef.afterClosed().subscribe(event => this.task.setEvent(event || this.task.event))
  }

}
