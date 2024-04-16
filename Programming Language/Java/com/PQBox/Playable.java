package com.PQBox;

// Declare the Playable interface
public interface Playable {
    // Declare the abstract method "play" that classes implementing this interface must provide
    void play();

    public default void score(){
        System.out.println("score default");
    }
} 
