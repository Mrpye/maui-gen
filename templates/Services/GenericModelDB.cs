using SQLite;

namespace {{.FullNS}} {

    //************************************************************
    //Generic database class to handle all the database operations
    //************************************************************
    public class GenericModelDB<M> where M : IEntity, new() {
        internal SQLiteAsyncConnection database;

        //**********************************************
        //Function to initialize the database connection
        //**********************************************
        public virtual async Task Init() {

            //*******************
            //Database connection
            //*******************
            if (database is not null) return;
            database = new SQLiteAsyncConnection(DBConstants.DatabasePath, DBConstants.Flags);

            //*****************
            //Create the tables
            //*****************
            try {
                {{range $field := .RootSchema.Schemas}}await database.CreateTableAsync<{{$field.FuncName}}Model>();
                {{end}}
            } catch (Exception e) {
                Debug.WriteLine(e);
            }


        }

    
        public async Task Close() {
            if(database!=null){
                await database.CloseAsync();
            }
            database = null;
        }


        //***********************************************
        //Function to get all the items from the database
        //***********************************************
        public async Task<List<M>> GetItemsAsync() {
            await Init();
            return await database.Table<M>().ToListAsync();
        }
        
        public async Task<List<M>> GetItemsWithChildrenAsync() {
            await Init();
            return await database.GetAllWithChildrenAsync<M>(recursive: true);
        }
 
 
        //***********************************************
        //Function to get all the items from the database
        //***********************************************
        public async Task<M> GetItemAsync(int id) {
            //****************************
            //Init the database connection
            //****************************
            await Init();

            //******************************
            //Get the item from the database
            //******************************
            return await database.Table<M>()
            .Where(i => i.Id == id)
            .FirstAsync();

        }

        //****************************************
        //Function to save an item to the database
        //****************************************
        public async Task<int> SaveItemAsync(M item) {

            //****************************
            //Init the database connection
            //****************************
            await Init();

            //*****************************
            //Save the item to the database
            //*****************************
            if (item.Id != 0)
                return await database.UpdateAsync(item);
            else
                return await database.InsertAsync(item);
        }

        public async Task SaveItemWithChildrenAsync(M item) {

            //****************************
            //Init the database connection
            //****************************
            await Init();

            //*****************************
            //Save the item to the database
            //*****************************
            //if (item.Id != 0)
                await database.InsertOrReplaceWithChildrenAsync(item, true);
            //else
            //    await database.InsertWithChildrenAsync(item, true);
        }



        public async Task<int> InsertItemsAsync(IEnumerable<M> item) {

            //****************************
            //Init the database connection
            //****************************
            await Init();

            return await database.InsertAllAsync(item, runInTransaction: true);

        }


        //********************************************
        //Function to delete an item from the database
        //********************************************
        public async Task<int> DeleteItemAsync(M item) {
            await Init();
            return await database.DeleteAsync(item);
        }

    }
}
