package com.medclever.pocketdoctor_free

import android.os.Bundle
import android.util.Log
import android.widget.ArrayAdapter
import android.widget.ListView
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat
import com.medclever.pocketdoctor_free.repositories.MessagesRepository
import com.medclever.pocketdoctor_free.utils.readFileFromAssets

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main)
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main)) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }

        val messagesRepo = MessagesRepository(assets, "data")
        val messages = messagesRepo.getAll().map { m -> m.message }
        val mainContent = findViewById<ListView>(R.id.mainContent)
        mainContent.adapter = ArrayAdapter(this, android.R.layout.simple_list_item_1, messages)
    }
}