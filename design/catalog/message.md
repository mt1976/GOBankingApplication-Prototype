# **Message** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Message** (message) |
|Endpoint 	    |**/Message...** [^1]|
|Endpoint Query |**Message**|
Glyph|**fas fa-info-circle** (text-info)
Friendly Name|**System Message**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Message/MessageList) [Exportable]
* **View** (/Message/MessageView)
* **Edit** (/Message/MessageEdit)
* **Save** (/Message/MessageSave)









##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **messageStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|false|true|false|false|||||Y|Id||false|false|false|text||
|**Message**|String|false|true|false|false|||||Y|Message||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/message_core.go |
| code | **dao** | /dao/message_core.go |
| code | **datamodel** | /datamodel/message_core.go |
| code | **menu** | /design/menu/message.json |
| html | **list** | /Message_List.html |
| html | **view** | /Message_View.html |
| html | **edit** | /Message_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:54**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field